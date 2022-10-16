package universe

import (
	"errors"

	"github.com/dragon-master-5892/universe/cachers"
	"github.com/dragon-master-5892/universe/loggers"
)

type ServiceInitParams struct {
	Broker *Broker
	Cacher *cachers.Cacher
	Logger *loggers.Logger
}

// Universe Service
type Service struct {
	name    string
	Broker  *Broker
	Actions map[string]*Action
}

func (s *Service) GetName() string {
	return s.name
}

func (s *Service) SetName(name string) (*Service, error) {
	if len(s.name) > 0 {
		return nil, errors.New("servce name setted before")
	}
	s.name = name
	return s, nil
}

func (s *Service) Init(initParams ServiceInitParams) *Service {
	if initParams.Broker != nil {
		s.Broker = initParams.Broker
	}

	if initParams.Cacher != nil {
		s.Broker.Cacher = initParams.Cacher
	}

	if initParams.Logger != nil {
		s.Broker.Logger = initParams.Logger
	}

	return s
}

func (s *Service) BrokerCall(action string, params map[string]interface{}) map[string]interface{} {
	return s.Broker.Call(action, params)
}

func (s *Service) RegisterAction(action *Action) *Service {
	s.Actions[action.Name] = action
	return s
}

func (s *Service) RegisterActions(actions ...*Action) *Service {
	for _, action := range actions {
		s.Actions[action.Name] = action
	}
	return s
}
