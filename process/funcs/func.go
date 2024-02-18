package funcs

import (
	"errors"
	"reflect"
	"strings"
)

var ErrFunc = errors.New("error function")

// 注册函数
var register = map[string]interface{}{
	"toLower": strings.ToLower,
	"toUpper": strings.ToUpper,
	"sum":     sum,
	"sqrt":    sqrt,
}

// 获取函数
func GetHandler(name string) (reflect.Value, error) {

	if f, ok := register[name]; ok {
		v := reflect.ValueOf(f)
		vt := reflect.TypeOf(f)
		//限制返回值有且只能为一个
		if v.Kind() == reflect.Func && vt.NumOut() == 1 {
			return v, nil
		}
	}
	return reflect.Value{}, ErrFunc
}
