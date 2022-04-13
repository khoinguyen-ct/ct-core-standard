package mongo_storage

import (
	"context"
	"github.com/ct-core-standard/config"
	"github.com/ct-core-standard/pkg/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const dbName = "mongodb"

var clients map[string]*mongo.Client
var logger = log.GetLogger()

func NewMongoStorage(ctx context.Context) (*mongo.Client, error) {
	if clients == nil {
		clients = make(map[string]*mongo.Client)
	}

	connectionStr := config.MongodbConfig().ConnectionString
	poolSize := config.MongodbConfig().PoolSize
	if client, ok := clients[connectionStr]; ok {
		if client.Ping(ctx, nil) == nil {
			return client, nil
		}
	}
	opts := options.Client().ApplyURI(connectionStr).SetMaxPoolSize(poolSize)
	timeoutCtx, _ := context.WithTimeout(ctx, time.Second * 10)
	client, err := mongo.Connect(timeoutCtx, opts)

	if err != nil {
		logger.Errorf("faile to connect mongodb: %v",err)
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Errorf("faile to connect mongodb: %v",err)
		return nil, err
	}
	logger.Info("Mongo connect succeeded")

	clients[connectionStr] = client

	return client, nil
}
