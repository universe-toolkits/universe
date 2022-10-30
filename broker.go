package universe

import (
	"github.com/universe-toolkits/universe/cachers"
	err "github.com/universe-toolkits/universe/errors"
	"github.com/universe-toolkits/universe/loggers"
	"github.com/universe-toolkits/universe/transporters"
)

type BrokerOptions struct {
}

type Broker struct {
	Namespace      string
	Id             string
	Transporter    transporters.Transporter
	Logger         loggers.Logger
	Cacher         cachers.Cacher
	Services       map[string]*Service
	RemoteServices map[string][]*UniverseNode
}

func (b *Broker) RegisterService(svc *Service) *Broker {
	b.Services[svc.GetName()] = svc
	return b
}

func (b *Broker) Call(action string, params map[string]interface{}) (map[string]interface{}, error) {

	if len(action) == 0 {
		return nil, err.InvalidActionNameError()
	}

	am, err := b.Transporter.Call("w", "e", params)
	return am, err
}

func (b *Broker) Start() *Broker {
	return b
}

func (b *Broker) Stop() *Broker {
	return b
}
