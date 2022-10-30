package universe

import (
	"github.com/universe-toolkits/universe/cachers"
	"github.com/universe-toolkits/universe/loggers"
)

type BrokerOptions struct {
}

type Broker struct {
	Transporter string
	Logger      *loggers.Logger
	Cacher      *cachers.Cacher
	Services    map[string]*Service
}

func (b *Broker) RegisterService(svc *Service) *Broker {
	b.Services[svc.GetName()] = svc
	return b
}

func (b *Broker) Call(action string, params map[string]interface{}) map[string]interface{} {
	return params
}
