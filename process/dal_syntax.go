package process

import (
	"dal/grammars/dal"
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/thoas/go-funk"
)

var AppVersions = []string{"api/1.0"}

// 语法检查
func Syntax(input string) error {

	fs := antlr.NewInputStream(input)

	lexer := dal.NewApiLexer(fs)

	ts := antlr.NewCommonTokenStream(lexer, 0)

	preParser := dal.NewApiParser(ts)

	errListener := NewErrorListener()

	//移除默认错误监听器
	preParser.RemoveErrorListeners()

	preParser.AddErrorListener(errListener)

	tree := preParser.Whole()

	antlr.ParseTreeWalkerDefault.Walk(NewSyntaxValidateTreeListener(), tree)

	return errListener.Err
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{}
}

// 语法监听器
type ErrorListener struct {
	Err error
	*antlr.DefaultErrorListener
}

func (d *ErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.Err = errors.New(fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func NewSyntaxValidateTreeListener() *SyntaxValidateTreeListener {
	return &SyntaxValidateTreeListener{}
}

type SyntaxValidateTreeListener struct {
	*dal.BaseApiListener
}

// 验证版本
func (s *SyntaxValidateTreeListener) EnterVersion(ctx *dal.VersionContext) {

	listener := ctx.GetParser().GetErrorListenerDispatch()
	if ctx.Version() == nil {
		listener.SyntaxError(ctx.GetParser(), nil, ctx.GetStart().GetLine(), ctx.GetStop().GetColumn(), "api version is empty", ctx.GetParser().GetError())
		return
	}
	text := ctx.Version().GetText()

	if !funk.ContainsString(AppVersions, text) {
		listener.SyntaxError(ctx.GetParser(), nil, ctx.GetStart().GetLine(), ctx.GetStop().GetColumn(), "api version is invalid", ctx.GetParser().GetError())
	}
}
