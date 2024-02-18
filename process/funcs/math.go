package funcs

import (
	"github.com/spf13/cast"
	"github.com/thoas/go-funk"
	"math"
)

func sum(v ...interface{}) float64 {
	return funk.Sum(v)
}
func sqrt(v interface{}) float64 {
	return math.Sqrt(cast.ToFloat64(v))
}
