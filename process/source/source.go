package source

import (
	"dal/grammars/source"
	"dal/process/driver"
	"dal/process/util"
	"errors"
	"github.com/antlr4-go/antlr/v4"
	"github.com/spf13/cast"
	"strings"
	"time"
)

// 默认超时时间 5秒
var DefaultTimeout = 5 * time.Second

// 默认尝试重连次数 0
var DefaultTry = 0

// 实例化资源
func NewSource(input string) *Source {

	is := antlr.NewInputStream(input)

	lexer := source.NewSourceLexer(is)

	ts := antlr.NewCommonTokenStream(lexer, 0)

	ps := source.NewSourceParser(ts)

	ps.RemoveErrorListeners()

	whole := ps.Whole()

	return &Source{Tree: whole}
}

// 源
type Source struct {
	Tree antlr.Tree
}

// 连接配置
func (s *Source) ConnConfig() (driver.ConnConfig, error) {

	cfg, err := s.parseDSN()
	if err != nil {
		return cfg, err
	}
	cfg.Timeout = s.Timeout()
	cfg.Try = s.Try()
	return cfg, nil
}

// 获取DSN
func (s *Source) DSN() string {
	return util.GetStrVal(s.Tree.(source.IWholeContext).Dsn().GetText())
}

// 获取Name
func (s *Source) Name() string {
	return s.Tree.(source.IWholeContext).Name().ID().GetText()
}

// 获取版本
func (s *Source) Version() string {
	return s.Tree.(source.IWholeContext).Version().GetText()
}

// 获取超时时间
func (s *Source) Timeout() time.Duration {
	timeout := s.Tree.(source.IWholeContext).Timeout()
	if timeout != nil {
		d, _ := time.ParseDuration(timeout.GetText() + "ms")
		return d
	}
	return DefaultTimeout
}

// 获取尝试重连次数
func (s *Source) Try() int {
	try := s.Tree.(source.IWholeContext).Try()
	if try != nil {
		return cast.ToInt(try.GetText())
	}
	return DefaultTry
}

// DSN 解析
func (s *Source) parseDSN() (driver.ConnConfig, error) {

	cfg := driver.ConnConfig{}

	dsn := s.DSN()
	ts := dsn
	driverType := dsn[:strings.IndexRune(dsn, ':')]
	if driverType == "" {
		return cfg, errors.New("dsn missing driver type")
	}
	cfg.Type = driverType

	ts = dsn[len(driverType)+len("://"):]

	var userPwd string

	if strings.IndexRune(dsn, '@') > -1 {
		userPwd = dsn[strings.Index(dsn, "//")+2 : strings.IndexRune(dsn, '@')]
	}

	if userPwd != "" {
		if strings.Index(userPwd, ":") > -1 {
			cfg.User = userPwd[:strings.Index(userPwd, ":")]
			cfg.Password = userPwd[strings.Index(userPwd, ":")+1:]
		} else {
			cfg.User = userPwd
		}
		ts = ts[len(userPwd)+1:]
	}
	net := ts[:strings.Index(ts, ":")]
	if net == "" {
		return cfg, errors.New("dsn missing network type")
	}
	cfg.Net = net

	ts = ts[len(net)+len("://"):]

	locPorts := ts[:strings.Index(ts, "/")]

	if locPorts == "" {
		return cfg, errors.New("dsn missing address")
	}
	cfg.LocPorts = make([]driver.LocPort, 0)
	locPortArr := strings.Split(locPorts, ",")
	for _, locPort := range locPortArr {
		loc := locPort[:strings.IndexRune(locPort, ':')]
		port := locPort[strings.IndexRune(locPort, ':')+1:]
		cfg.LocPorts = append(cfg.LocPorts, driver.LocPort{
			Loc:  loc,
			Port: cast.ToInt(port),
		})
	}
	ts = ts[len(locPorts):]
	var ps string
	if strings.IndexRune(ts, '?') != -1 {
		ps = ts[strings.IndexRune(ts, '?')+1:]
	}

	sd := ts
	if ps != "" {
		sd = ts[:strings.IndexRune(ts, '?')]
	}
	if strings.Count(sd, "/") == 1 {
		cfg.Database = ts[1:]
	} else if strings.Count(ts, "/") == 2 {
		cfg.Schema = ts[1:strings.LastIndex(ts, "/")]
		cfg.Database = ts[strings.LastIndex(ts, "/")+1:]
	} else {
		return cfg, errors.New("dsn missing schema or database")
	}

	params := make([]string, 0)
	if ps != "" {
		params = strings.Split(ps, "&")
	}

	cfg.Params = map[string]string{}
	for _, p := range params {
		options := strings.Split(p, "=")
		if len(options) != 2 {
			return cfg, errors.New("error dsn options")
		}
		cfg.Params[options[0]] = options[1]
	}

	return cfg, nil
}
