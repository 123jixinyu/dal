package util

import (
	"github.com/spf13/cast"
	"reflect"
	"regexp"
	"strings"
)

// 获取字符串实际值
func GetStrVal(str string) string {
	return str[strings.Index(str, "'")+1 : strings.LastIndex(str, "'")]
}

// 转换数组
func GetSlice(str string) []interface{} {

	arr := make([]interface{}, 0)

	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.TrimSpace(str)
	if str == "" {
		return arr
	}
	strArr := strings.Split(str, ",")

	for _, v := range strArr {

		if ok, _ := regexp.MatchString("\\'.*\\'", v); ok {

			arr = append(arr, GetStrVal(v))
		}
		if v == "true" || v == "false" {
			arr = append(arr, cast.ToBool(v))
		}
		if ok, _ := regexp.MatchString("^(-?\\d+)(\\.\\d+)?$ 或 ^-?([1-9]\\d*\\.\\d*|0\\.\\d*[1-9]\\d*|0?\\.0+|0)$", v); ok {
			arr = append(arr, cast.ToFloat64(v))
		}
		arr = append(arr, cast.ToInt(v))
	}
	return arr
}

// 转换类型

func Convert(v interface{}, t reflect.Type) (interface{}, error) {

	var sv = cast.ToString(v)
	var nv interface{} = sv

	switch t.String() {
	case "int":
		nv = cast.ToInt(sv)
	case "int8":
		nv = cast.ToInt8(sv)
	case "int16":
		nv = cast.ToInt16(sv)
	case "int32":
		nv = cast.ToInt32(sv)
	case "int64":
		nv = cast.ToInt64(sv)
	case "sql.RawBytes":
		nv = cast.ToString(sv)
	case "sql.NullTime":
		if v == nil {
			return nil, nil
		}
	case "sql.NullString":
		if v == nil {
			return nil, nil
		}
	case "sql.NullInt16":
		if v == nil {
			return nil, nil
		}

		nv = cast.ToInt16(string(v.([]uint8)))
	case "sql.NullInt32":
		if v == nil {
			return nil, nil
		}
		nv = cast.ToInt32(string(v.([]uint8)))
	case "sql.NullInt64":
		if v == nil {
			return nil, nil
		}

		nv = cast.ToInt64(string(v.([]uint8)))
	}
	if v != nil {
		//fmt.Println(fmt.Sprintf("origin type: %s, origin value:%v", reflect.ValueOf(v).Type().String(), v), "New Type:"+t.String(), "New Value:", nv)
	}
	return nv, nil
}
