package universe

import (
	"github.com/dragon-master-5892/universe/cachers"
	"github.com/dragon-master-5892/universe/loggers"
)

type Context struct {
	Params map[string]interface{}
	Meta   map[string]interface{}
	Broker *Broker
	Cacher *cachers.Cacher
	Logger *loggers.Logger
}

func (ctx *Context) GetParam(key string) interface{} {
	return ctx.Params[key]
}

func (ctx *Context) GetMetaItem(key string) interface{} {
	return ctx.Meta[key]
}
