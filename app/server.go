package app

import (
	"context"
	"fmt"
	"github.com/ct-core-standard/internal/usecase"
	"github.com/ct-core-standard/pkg/log"
	svr "github.com/ct-core-standard/pkg/protobuf"
	"github.com/ct-core-standard/pkg/service_container"
	"github.com/facebookgo/inject"
)

var logger = log.GetLogger()

type serverAPI struct {
	Repos *usecase.Repos `inject:""`
}

func NewAdListingServiceServer(ctx context.Context) svr.AdListingServiceServer {
	var s serverAPI
	services := service_container.GetServices()
	if err := s.injectDependency(services); err != nil {
		logger.Errorf("build service graph failed: %v", err)
		return &s
	}
	for _, service := range services {
		if err := service.Instance.Init(ctx); err != nil {
			logger.Errorf("service: %s init failed", service.Name)
			return &s
		}
		//logger.Infof("service: %s init succeed", service.Name)
	}

	return &s
}

func (s *serverAPI) GetAdInfo(ctx context.Context, req *svr.GetAdInfoRequest) (res *svr.GetAdInfoResponse, err error) {
	uc := usecase.NewAdListingUC(s.Repos)
	data, err := uc.GetByListID(ctx, req.ListId)
	if err != nil {
		return nil, err
	}
	res = &data
	return res, nil
}

func (s *serverAPI) injectDependency(services []*service_container.Component) error {
	objs := []interface{}{
		s,
	}

	for _, service := range services {
		objs = append(objs, service.Instance)
	}

	var serviceGraph inject.Graph

	for _, obj := range objs {
		if err := serviceGraph.Provide(&inject.Object{Value: obj}); err != nil {
			return fmt.Errorf("%s: %s", "failed to provide object to the graph", err)
		}
	}

	if err := serviceGraph.Populate(); err != nil {
		return fmt.Errorf("%s: %s", "failed to populate service dependency", err)
	}

	return nil
}
