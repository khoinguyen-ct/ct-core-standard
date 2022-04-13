package mongo_storage

import (
	"context"
	"github.com/ct-core-standard/internal/model"
	"github.com/ct-core-standard/pkg/service_container"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAdListing interface {
	MongoGetByAdID(ctx context.Context, adID int64) (model.AdListing, error)
}

type mongoAdListingImpl struct {
	mongoClient *mongo.Client
}

func init() {
	service_container.RegisterService(&mongoAdListingImpl{})
}

func(m *mongoAdListingImpl) Init(ctx context.Context) error {
	client, err := NewMongoStorage(ctx)
	if err != nil {
		logger.Errorf("failed to get mongo client")
	}
	m.mongoClient = client
	return err
}

func (m *mongoAdListingImpl) MongoGetByAdID(ctx context.Context, adID int64) (model.AdListing, error) {
	return model.AdListing{}, nil
}
