package driver

import (
	"fmt"
	"time"
)

var register = map[string]Driver{
	"mysql": &MySQL{},
}

type Driver interface {
	SetConfig(cfg ConnConfig)
	Exec(sql string) ([]map[string]interface{}, error)
}

// New driver
func NewDriver(cfg ConnConfig) Driver {

	dr, ok := register[cfg.Type]
	if !ok {
		panic(fmt.Sprintf("cannot support driver type %s", cfg.Type))
	}
	dr.SetConfig(cfg)
	return dr
}

// 基础
type BaseDriver struct {
	ConnConfig ConnConfig
}

func (b *BaseDriver) SetConfig(cfg ConnConfig) {
	b.ConnConfig = cfg
}

type ConnConfig struct {
	Type     string            `json:"type"`
	User     string            `json:"user"`
	Password string            `json:"password"`
	Net      string            `json:"net"`
	Schema   string            `json:"schema"`
	Database string            `json:"database"`
	LocPorts []LocPort         `json:"loc_ports"`
	Params   map[string]string `json:"params"`
	Timeout  time.Duration     `json:"timeout"`
	Try      int               `json:"try"`
}

type LocPort struct {
	Loc  string `json:"loc"`
	Port int    `json:"port"`
}
