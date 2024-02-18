package process

import (
	"bytes"
	"dal/grammars/dal"
	"dal/process/driver"
	"dal/process/funcs"
	"dal/process/plugin"
	"dal/process/source"
	"dal/process/util"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"
)

const VarPrefix = "$"

var SuccessCode = 1
var SuccessMessage = "success"
var ErrorCode = -1
var UnKnownError = errors.New("unknown error")
var InvalidOperator = errors.New("invalid operator")

// 初始化Dal
func NewDal(input string) *Dal {
	d := &Dal{
		Validator: validator.New(),
		Vars:      map[string]Var{},
	}

	d.Tree = d.parser(input)

	d.Name = d.Tree.Name().GetText()
	d.Version = d.Tree.Version().GetText()

	return d
}

type IApi interface {
	Method() string
	Route() string
	Handler() echo.HandlerFunc
	MiddlewareFunc() []echo.MiddlewareFunc
}

// DAL核心
type Dal struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Tree      dal.IWholeContext
	Ctx       echo.Context
	Validator *validator.Validate
	Vars      map[string]Var `json:"vars"`
}

// 变量
type Var struct {
	Alias string       `json:"alias"`
	Value interface{}  `json:"value"`
	Kind  reflect.Kind `json:"kind"`
}

// 获取请求方式
func (d *Dal) Method() string {
	return d.Tree.Endpoint().Method().GetText()
}

// 获取路由
func (d *Dal) Route() string {
	return d.Tree.Endpoint().Route().GetText()
}

// api处理逻辑
func (d *Dal) Handler() echo.HandlerFunc {

	return func(ctx echo.Context) error {
		d.Ctx = ctx

		//解析请求参数，注入请求变量
		if err := d.resolveRequestParams(); err != nil {
			return d.response(nil, err)
		}

		//验证请求参数
		if err := d.validate(); err != nil {
			fmt.Println(err)
			return d.response(nil, err)
		}

		//执行各个阶段
		if err := d.stageExecution(); err != nil {
			return d.response(nil, err)
		}
		//获取响应内容
		content, err := d.getContent()
		if err != nil {
			return d.response(nil, err)
		}

		return d.response(content, nil)
	}
}

// api中间件
func (d *Dal) MiddlewareFunc() []echo.MiddlewareFunc {

	fs := make([]echo.MiddlewareFunc, 0)

	allPlugins := d.Tree.AllPlugin()

	for _, v := range allPlugins {
		pluginName := v.Plugin().GetText()
		newPlugin := plugin.New(pluginName)
		if v.PluginParams() != nil {
			for _, p := range v.PluginParams().AllPluginParam() {
				flag := p.ID().GetText()
				value := p.String_().GetText()
				newPlugin.Set(flag, value)
			}
		}
		fs = append(fs, newPlugin.Func())
	}

	return fs
}

// 处理
func (d *Dal) stageExecution() error {

	stages := d.Tree.AllStage()

	for _, s := range stages {

		stageVarPrefix := fmt.Sprintf("stage.%s", s.ID().GetText())

		handles := s.Handles().AllHandle()

		for _, h := range handles {
			stageVarName := h.ID().GetText()
			var value interface{}
			if h.Value() != nil {
				data, err := d.value(h.Value(), stageVarPrefix)
				if err != nil {
					return err
				}
				value = data
			}
			if h.Query() != nil {
				data, err := d.query(h.Query(), stageVarPrefix)
				if err != nil {
					return err
				}
				value = data
			}
			if h.Express() != nil {
				r, err := d.express(h.Express(), stageVarPrefix)
				if err != nil {
					return err
				}
				value = r
			}
			d.addProcessVar(stageVarName, value, stageVarPrefix)
		}
	}
	return nil
}

// 获取响应内容
func (d *Dal) getContent() (interface{}, error) {

	content := d.Tree.Return_().Content()

	if content.Body().Express() != nil {
		return d.express(content.Body().Express(), "")
	}

	if content.Body().Object() != nil {
		var data = map[string]interface{}{}
		if err := d.parseObject(content.Body().Object(), &data); err != nil {
			return nil, err
		}
		return data, nil
	}

	return nil, errors.New("content is invalid")
}

// 解析object
func (d *Dal) parseObject(ctx dal.IObjectContext, data *map[string]interface{}) error {

	pairs := ctx.AllObjectPair()

	mp := *data

	for _, p := range pairs {

		if p.ObjectValue().Value() != nil {
			ve, err := d.value(p.ObjectValue().Value(), "")
			if err != nil {
				return err
			}
			mp[p.ID().GetText()] = ve
		}
		if p.ObjectValue().Express() != nil {
			v, err := d.express(p.ObjectValue().Express(), "")
			if err != nil {
				return err
			}
			mp[p.ID().GetText()] = v
		}
		if p.ObjectValue().Object() != nil {
			objectData := map[string]interface{}{}
			if err := d.parseObject(p.ObjectValue().Object(), &objectData); err != nil {
				return err
			}
			if _, ok := mp[p.ID().GetText()]; !ok {
				mp[p.ID().GetText()] = map[string]interface{}{}
			}
			mp[p.ID().GetText()] = objectData
		}
	}
	return nil
}

// 响应
func (d *Dal) response(data interface{}, err error) error {

	responseCode, err := d.getResponseCode()
	if err != nil {
		return err
	}
	d.Ctx.Response().Status = responseCode

	for k, arr := range d.getResponseHeader() {
		for _, v := range arr {
			d.Ctx.Response().Header().Add(k, v)
		}
	}
	body := map[string]interface{}{}

	var code = SuccessCode
	var msg = SuccessMessage
	if err != nil {
		code = ErrorCode
		msg = err.Error()
	}
	body["code"] = code
	body["msg"] = msg
	body["data"] = data

	responseContentType := d.getResponseHeader().Get(echo.HeaderContentType)

	buf := make([]byte, 0)

	if responseContentType == echo.MIMEApplicationJSON {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return err
		}
		buf = jsonData
	}

	if responseContentType == echo.MIMEApplicationXML {
		xmlData, err := xml.Marshal(body)
		if err != nil {
			return err
		}
		buf = xmlData
	}

	_, err = d.Ctx.Response().Write(buf)
	if err != nil {
		return err
	}
	d.Ctx.Set("response_body", string(buf))
	return nil
}

// 获取响应头
func (d *Dal) getResponseHeader() http.Header {

	hd := http.Header{}

	if d.Tree.Return_().ResponseHeaders() != nil {
		responseHeaders := d.Tree.Return_().ResponseHeaders().AllResponseHeader()

		for _, v := range responseHeaders {
			headerK, headerV := v.String_(0).GetText(), v.String_(1).GetText()
			hd.Add(headerK, headerV)
		}
	}

	//默认application/json响应头
	if hd.Get(echo.HeaderContentType) == "" {
		hd.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}
	return hd
}

// 获取响应码
func (d *Dal) getResponseCode() (int, error) {

	rc := d.Tree.Return_().ResponseCode()
	if rc == nil || rc.GetText() == "" {
		//默认200
		return http.StatusOK, nil
	}
	if rc.Int() != nil {
		return cast.ToInt(rc.Int().GetText()), nil
	}
	if rc.Express() != nil {
		v, err := d.express(rc.Express(), "")
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return cast.ToInt(v), nil
	}

	return http.StatusOK, nil
}

// 是否有preflight请求
func (d *Dal) HasPreflight() bool {
	preflights := d.Tree.AllPreflight()
	return len(preflights) > 0
}

// 获取preflight处理
func (d *Dal) GetPreflightHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		for _, preflight := range d.Tree.AllPreflight() {
			c.Response().Header().Add(preflight.PreflightHeader().GetText(), util.GetStrVal(preflight.String_().GetText()))
		}
		return nil
	}
}

// 参数验证
func (d *Dal) validate() error {

	v := d.Tree.Validate()
	if v == nil {
		return nil
	}

	validates := v.FieldValidates().AllFieldValidate()

	for _, v := range validates {

		requestName := v.Var().GetText()
		validateRule := util.GetStrVal(v.String_().GetText())
		requestVal := d.getVar(requestName).Value

		err := d.validateVar(requestName, requestVal, validateRule)
		if err != nil {
			return err
		}
	}
	return nil
}

// dsl解析
func (d *Dal) parser(input string) dal.IWholeContext {

	fs := antlr.NewInputStream(input)

	lexer := dal.NewApiLexer(fs)

	ts := antlr.NewCommonTokenStream(lexer, 0)

	parser := dal.NewApiParser(ts)

	parser.RemoveErrorListeners()

	return parser.Whole()
}

// 处理请求参数
func (d *Dal) resolveRequestParams() error {

	d.addProcessVar("url", d.Ctx.Request().RequestURI, "request")
	d.addProcessVar("ip", d.Ctx.RealIP(), "request")
	d.addProcessVar("remote_addr", d.Ctx.Request().RemoteAddr, "request")
	d.addProcessVar("host", d.Ctx.Request().Host, "request")
	d.addProcessVar("path", d.Ctx.Path(), "request")
	d.addProcessVar("is_tls", d.Ctx.IsTLS(), "request")
	d.addProcessVar("query_string", d.Ctx.QueryString(), "request")

	for name, value := range d.Ctx.Request().Header {
		name = strings.ToLower(name)
		if len(value) == 1 {
			d.addProcessVar(name, value[0], "request.header")
			continue
		}
		d.addProcessVar(name, value, "request.header")
	}

	for queryParamKey, queryParamVal := range d.Ctx.QueryParams() {
		if len(queryParamVal) == 1 {
			d.addProcessVar(queryParamKey, queryParamVal[0], "request.url")
			continue
		}
		d.addProcessVar(queryParamKey, queryParamVal, "request.url")
	}
	err := d.Ctx.Request().ParseForm()
	if err != nil {
		return err
	}
	contentType := strings.ToLower(d.Ctx.Request().Header.Get("Content-Type"))

	formValues := map[string][]string{}
	if strings.HasPrefix(contentType, echo.MIMEMultipartForm) {
		if err := d.Ctx.Request().ParseMultipartForm(1024 * 1024); err != nil {
			return err
		}
		formValues = d.Ctx.Request().MultipartForm.Value
	}
	if contentType == echo.MIMEApplicationForm {
		formValues = d.Ctx.Request().PostForm
	}
	for formParamKey, formParamVal := range formValues {
		if len(formParamVal) == 1 {
			d.addProcessVar(formParamKey, formParamVal[0], "request.form")
			continue
		}
		d.addProcessVar(formParamKey, formParamVal, "request.form")
	}
	if contentType == echo.MIMEApplicationJSON {
		body, err := io.ReadAll(d.Ctx.Request().Body)
		if err != nil {
			return err
		}
		d.Ctx.Request().Body = io.NopCloser(bytes.NewBuffer(body))
		d.addProcessVar("json", string(body), "request")
	}

	return nil
}

// 获取变量
func (d *Dal) getVarE(alias string, group string) (Var, error) {

	//局部变量
	if !strings.Contains(alias, ".") {
		if group == "" {
			return Var{}, errors.New("group is required when getting variable")
		}
		alias = strings.Replace(alias, VarPrefix, VarPrefix+group+".", -1)
	}

	if v, ok := d.Vars[alias]; ok {
		return v, nil
	}

	jsonPrefix := VarPrefix + "request.json."
	if strings.HasPrefix(alias, jsonPrefix) {
		jsonSearch := alias[len(jsonPrefix):]
		jsonVal := d.Vars[VarPrefix+"request.json"].Value.(string)
		result := gjson.Get(jsonVal, jsonSearch)
		if !result.Exists() {
			return Var{}, errors.New(fmt.Sprintf("can not found the var:%s", alias))
		}
		return Var{
			Value: result.Value(),
			Alias: alias,
			Kind:  reflect.TypeOf(result.Value()).Kind(),
		}, nil
	}
	return Var{}, errors.New(fmt.Sprintf("can not found the var:%s", alias))
}

// 获取变量 无视错误
func (d *Dal) getVar(alias string) Var {
	v, _ := d.getVarE(alias, "")
	return v
}

// 在处理过程中保存变量
func (d *Dal) addProcessVar(name string, value interface{}, group string) {

	name = strings.TrimSpace(name)

	if name == "" {
		log.Fatal("var name can not be empty")
		return
	}
	var kind = reflect.Invalid
	if value != nil {
		kind = reflect.TypeOf(value).Kind()
	}

	alias := VarPrefix + group + "." + name

	d.Vars[alias] = Var{
		Alias: alias,
		Value: value,
		Kind:  kind,
	}
}

// 获取变量
func (d *Dal) getVarString() string {

	ss := make([]string, 0)

	for k, v := range d.Vars {
		ss = append(ss, fmt.Sprintf("key:%s,value:%v，kind: %s", k, v.Value, v.Kind.String()))
	}

	return strings.Join(ss, "\n")
}

// 变量验证
func (d *Dal) validateVar(name string, val interface{}, rule string) error {

	err := d.Validator.Var(val, rule)
	if err == nil {
		return nil
	}
	return errors.New(strings.ReplaceAll(err.Error(), "''", "'"+name+"'"))
}

// 处理表达式
func (d *Dal) express(e dal.IExpressContext, group string) (interface{}, error) {
	return NewExpress(d, group).Walk(e.Calculate())
}

// 处理值
func (d *Dal) value(v dal.IValueContext, group string) (interface{}, error) {

	var value interface{}

	if v.Int() != nil {
		value = cast.ToInt(v.Int().GetText())
	}
	if v.String_() != nil {
		value = util.GetStrVal(v.String_().GetText())
	}
	if v.Bool() != nil {
		value = cast.ToBool(v.Bool().GetText())
	}
	if v.Float() != nil {
		value = cast.ToFloat64(v.Float().GetText())
	}
	if v.Array() != nil {
		value = util.GetSlice(v.Array().GetText())
	}
	if v.Var() != nil {
		data, err := d.getVarE(v.Var().GetText(), group)
		if err != nil {
			return nil, err
		}
		value = data.Value
	}
	return value, nil
}

// 处理多个值
func (d *Dal) values(vs []dal.IValueContext, group string) ([]interface{}, error) {

	values := make([]interface{}, 0)
	if len(vs) > 0 {
		for _, value := range vs {
			ve, err := d.value(value, group)
			if err != nil {
				return values, err
			}
			values = append(values, ve)
		}
	}
	return values, nil
}

// 获取query字符串
func (d *Dal) query(v dal.IQueryContext, group string) (interface{}, error) {

	sourceAlias := v.Var().GetText()

	src, err := source.DefaultStore.Find(sourceAlias)
	if err != nil {
		return nil, err
	}
	sql, err := d.getSql(v, group)
	if err != nil {
		return nil, err
	}
	cfg, err := src.ConnConfig()
	if err != nil {
		return nil, err
	}
	return driver.NewDriver(cfg).Exec(sql)
}

// 获取sql
func (d *Dal) getSql(v dal.IQueryContext, group string) (string, error) {

	template, err := d.getTemplate(v.Sql(), group)
	if err != nil {
		return "", err
	}
	values, err := d.values(v.AllValue(), group)
	if err != nil {
		return "", err
	}
	return d.buildSql(template, values)
}

// 获取模版
func (d *Dal) getTemplate(ctx dal.ISqlContext, group string) (string, error) {

	var template string

	if ctx.Var() != nil {
		processVar, err := d.getVarE(ctx.Var().GetText(), group)
		if err != nil {
			return "", err
		}
		if processVar.Kind != reflect.String {
			return "", errors.New("invalid sql type")
		}
		template = processVar.Value.(string)
	}
	if ctx.String_() != nil {
		template = util.GetStrVal(ctx.String_().GetText())
	}
	return template, nil
}

// build sql
func (d *Dal) buildSql(template string, values []interface{}) (string, error) {

	if strings.Count(template, "?") != len(values) {
		return "", errors.New("invalid sql parameter numbers")
	}

	ss := make([]any, 0)
	for _, v := range values {
		var s string
		rv := reflect.ValueOf(v)

		switch rv.Kind() {
		case reflect.Slice:
			tpArr := make([]string, 0)
			for i := 0; i < rv.Len(); i++ {
				if rv.Index(i).Kind() == reflect.String {
					tpArr = append(tpArr, "'"+rv.Index(i).String()+"'")
				} else {
					tpArr = append(tpArr, cast.ToString(rv.Index(i).Interface()))
				}
			}
			s = "(" + strings.Join(tpArr, ",") + ")"
		case reflect.String:
			s = "'" + v.(string) + "'"
		default:
			s = cast.ToString(v)
		}
		ss = append(ss, s)
	}
	ts := strings.ReplaceAll(template, "?", "%s")
	v := fmt.Sprintf(ts, ss...)

	return v, nil
}

func NewExpress(d *Dal, group string) *Express {
	return &Express{Dal: d, Group: group}
}

type Express struct {
	Dal   *Dal
	Group string
}

// 遍历表达式
func (e *Express) Walk(calculate dal.ICalculateContext) (interface{}, error) {

	if calculate.LeftParenthesis() != nil {
		return e.Walk(calculate.Calculate(0))
	}

	if calculate.MuOperator() != nil {

		left, err := e.WalkParam(calculate.Param())
		if err != nil {
			return nil, err
		}
		right, err := e.Walk(calculate.Calculate(0))
		if err != nil {
			return nil, err
		}
		return e.Operator(left, calculate.MuOperator().GetText(), right)
	}
	if calculate.ArOperator() != nil {
		left, err := e.Walk(calculate.Calculate(0))
		if err != nil {
			return nil, err
		}

		right, err := e.Walk(calculate.Calculate(1))
		if err != nil {
			return nil, err
		}
		return e.Operator(left, calculate.ArOperator().GetText(), right)
	}

	if calculate.MetaExpress() != nil {
		return e.WalkMetaExpress(calculate.MetaExpress())
	}
	if calculate.Param() != nil {
		return e.WalkParam(calculate.Param())
	}
	if calculate.Func_() != nil {
		return e.WalkFunc(calculate.Func_())
	}
	return nil, errors.New(fmt.Sprintf("unknown calculate:%v", calculate.GetText()))
}

// 遍历参数
func (e *Express) WalkParam(ctx dal.IParamContext) (interface{}, error) {

	children := ctx.GetChildren()

	for _, child := range children {

		switch child.(type) {
		case *dal.FuncContext:
			return e.WalkFunc(child.(*dal.FuncContext))
		case *dal.ValueContext:
			ve, err := e.Dal.value(child.(*dal.ValueContext), e.Group)
			if err != nil {
				return nil, err
			}
			return ve, nil
		default:
			log.Fatal("Unknown", reflect.TypeOf(child).String())
		}
	}
	return nil, nil
}

// 遍历函数
func (e *Express) WalkFunc(ctx dal.IFuncContext) (interface{}, error) {

	funcName := ctx.ID().GetText()
	params := ctx.Params().AllParam()
	paramValues := make([]interface{}, 0)
	for _, param := range params {
		paramValue, err := e.WalkParam(param)
		if err != nil {
			return paramValue, err
		}
		paramValues = append(paramValues, paramValue)
	}
	return e.CallFunc(funcName, paramValues)
}

// 遍历三元表达式
func (e *Express) WalkMetaExpress(ctx dal.IMetaExpressContext) (interface{}, error) {

	pv1, err := e.WalkParam(ctx.Param(0))
	if err != nil {
		return nil, err
	}

	pv2, err := e.WalkParam(ctx.Param(1))
	if err != nil {
		return nil, err
	}

	if ctx.JudgeOperator().GetText() == "!=" {
		if pv1 != pv2 {
			return e.Walk(ctx.Calculate(0))
		}
		return e.Walk(ctx.Calculate(1))
	}

	if ctx.JudgeOperator().GetText() == "==" {
		if pv1 == pv2 {
			return e.Walk(ctx.Calculate(0))
		}
		return e.Walk(ctx.Calculate(1))
	}

	return nil, errors.New("invalid meta expression")
}

// 调用函数
func (e *Express) CallFunc(funcName string, params []interface{}) (interface{}, error) {

	handler, err := funcs.GetHandler(funcName)
	if err != nil {
		return nil, err
	}
	values := make([]reflect.Value, 0)

	for _, p := range params {
		values = append(values, reflect.ValueOf(p))
	}

	val := handler.Call(values)[0].Interface()

	return val, nil
}

// 常规运算
func (e *Express) Operator(left interface{}, operator string, right interface{}) (interface{}, error) {

	leftF := cast.ToFloat64(left)
	rightF := cast.ToFloat64(right)

	switch operator {
	case "+":
		return leftF + rightF, nil
	case "-":
		return leftF - rightF, nil
	case "*":
		return leftF * rightF, nil
	case "/":
		return leftF / rightF, nil
	}
	return nil, InvalidOperator
}
