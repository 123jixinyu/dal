package plugin

import (
	"github.com/labstack/echo/v4"
	"sync"
)

var AllPlugins = []IPlugin{
	&Log{},
}

type IPlugin interface {
	ID() string
	Func() echo.MiddlewareFunc
	Set(key string, value string)
	Get(key string) string
}

func New(id string) IPlugin {

	for _, plugin := range AllPlugins {
		if plugin.ID() == id {
			return plugin
		}
	}
	panic("no such plugin")
}

type Params struct {
	Map sync.Map
}

func (p *Params) Set(key string, value string) {
	p.Map.Delete(key)
	p.Map.Store(key, value)
}
func (p *Params) Get(key string) string {

	v, _ := p.Map.Load(key)
	if v == nil {
		return ""
	}
	return v.(string)
}
