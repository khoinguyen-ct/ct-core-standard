package es_storage

import (
	"context"
	"github.com/ct-core-standard/config"
	"github.com/ct-core-standard/pkg/log"
	esv7 "github.com/elastic/go-elasticsearch/v7"
	"strings"
)

var logger = log.GetLogger()

func NewEsClient(ctx context.Context) (*esv7.Client, error) {
	adders := strings.Split(config.ElasticsearchConfig().Addr,",")
	cfg := esv7.Config{
		Addresses: adders,
	}
	esClient, err := esv7.NewClient(cfg)
	if err != nil {
		logger.Errorf("failed to init es client, err: %v", err)
	}
	return esClient, err
}
