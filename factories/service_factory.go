package factories

import (
	uni "github.com/dragon-master-5892/universe"
)

func NewService(serviceName string) *uni.Service {
	svc := &uni.Service{}
	svc.SetName(serviceName)
	return svc
}
