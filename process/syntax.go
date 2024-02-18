package process

import (
	"dal/grammars/version"
	"errors"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

// 获取类型
func GetKindVersion(input string) (string, error) {

	fs := antlr.NewInputStream(input)

	lexer := version.NewVersionLexer(fs)

	ts := antlr.NewCommonTokenStream(lexer, 0)

	preParser := version.NewVersionParser(ts)

	errListener := NewKindErrorListener()

	//移除默认错误监听器
	preParser.RemoveErrorListeners()

	preParser.AddErrorListener(errListener)

	tree := preParser.Whole()

	v := tree.Version().Version().GetText()

	return v, errListener.Err
}

func NewKindErrorListener() *ErrorListener {
	return &ErrorListener{}
}

// 语法监听器
type KindErrorListener struct {
	Err error
	*antlr.DefaultErrorListener
}

func (d *KindErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, e antlr.RecognitionException) {
	d.Err = errors.New(fmt.Sprintf("line %d:%d %s", line, column, msg))
}
