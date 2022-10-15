package universe

import (
	"errors"
)

type ServiceInitParams struct {
	broker *Broker
}

// Universe Service
type Service struct {
	name   string
	broker *Broker
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
	if initParams.broker != nil {
		s.broker = initParams.broker
	}
	return s
}

func (s *Service) BrokerCall(action string, params interface{}) interface{} {
	return s.broker.Call(action, params)
}
