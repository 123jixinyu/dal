package plugin

import (
	"dal/process/util"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Log struct {
	Params
}

func (*Log) ID() string {
	return "log"
}

func (l *Log) Func() echo.MiddlewareFunc {

	loggerConfig := middleware.LoggerConfig{
		Skipper:          middleware.DefaultSkipper,
		Format:           l.Format() + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}

	return middleware.LoggerWithConfig(loggerConfig)
}

// 日志格式
func (l *Log) Format() string {

	defaultFormat := `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
		`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
		`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
		`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}`

	format := l.Get("format")
	if format == "" {
		return defaultFormat
	}
	return util.GetStrVal(format)
}
