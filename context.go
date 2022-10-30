package universe

import (
	"github.com/universe-toolkits/universe/cachers"
	"github.com/universe-toolkits/universe/loggers"
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
