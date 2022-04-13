package service_container

import (
	"context"
	"reflect"
)

type Service interface {
	Init(ctx context.Context) error
}

type Component struct {
	Name         string
	Instance     Service
}

var services []*Component

func RegisterService(service Service){
	services = append(services, &Component{
		Name:     reflect.TypeOf(service).Elem().Name(),
		Instance: service,
	})
}

func GetServices() []*Component {
	return services
}