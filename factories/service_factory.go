package factories

import (
	uni "github.com/universe-toolkits/universe"
)

func NewService(serviceName string) *uni.Service {
	svc := &uni.Service{}
	svc.SetName(serviceName)
	return svc
}

func NewAction(actionName string) *uni.Action {
	return &uni.Action{
		Name: actionName,
	}
}
