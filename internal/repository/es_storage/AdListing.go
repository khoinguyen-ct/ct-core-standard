package es_storage

import (
	"context"
	"github.com/ct-core-standard/internal/model"
	"github.com/ct-core-standard/pkg/service_container"
	esv7 "github.com/elastic/go-elasticsearch/v7"
)

type ESAdListing interface {
	ESGetByAdID(ctx context.Context, adID int64) (model.AdListing, error)
}

type esAdListingImpl struct {
	esClient *esv7.Client
}

func init() {
	service_container.RegisterService(&esAdListingImpl{})
}

func(es *esAdListingImpl) Init(ctx context.Context) error {
	client, err := NewEsClient(ctx)
	if err != nil {
		logger.Errorf("failed to get elasticsearch client")
	}
	es.esClient = client
	return err
}

func (es *esAdListingImpl) ESGetByAdID(ctx context.Context, adID int64) (model.AdListing, error) {
	return model.AdListing{}, nil
}
