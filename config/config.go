package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server        `mapstructure:"server"`
	Service       `mapstructure:"service"`
	Mongodb       `mapstructure:"mongodb"`
	Elasticsearch `mapstructure:"elasticsearch"`
	Postgresql    `mapstructure:"postgresql"`
	Kafka         `mapstructure:"kafka"`
	Rabbitmq      `mapstructure:"rabbitmq"`
	Redis         `mapstructure:"redis"`
}

type Server struct {
	HTTPPort string `mapstructure:"http_port"`
	GRPCPort string `mapstructure:"grpc_port"`
	LogLevel string `mapstructure:"log_level"`
}

type Service struct {
	UserAdsDomain        string `mapstructure:"user_ads_domain"`
	AdListingDomain      string `mapstructure:"ad_listing_domain"`
	SpineDomain          string `mapstructure:"spine_domain"`
	AccessTradeDomain    string `mapstructure:"access_trade_domain"`
	SchemaRegistryDomain string `mapstructure:"schema_registry_domain"`
}

type Mongodb struct {
	ConnectionString string `mapstructure:"connection_string"`
	PoolSize         uint64 `mapstructure:"pool_size"`
}

type Elasticsearch struct {
	Addr      string `mapstructure:"addr"`
	Index     string `mapstructure:"index"`
	IndexType string `mapstructure:"index_type"`
}

type Postgresql struct {
	Addr      string `mapstructure:"addr"`
	Index     string `mapstructure:"index"`
	IndexType string `mapstructure:"index_type"`
}

type Kafka struct {
	Brokers           string `mapstructure:"brokers"`
	ConsumerGroup     string `mapstructure:"consumer_group"`
	TopicAds          string `mapstructure:"topic_ads"`
	TopicActionStates string `mapstructure:"topic_action_states"`
	InitOffset        string `mapstructure:"init_offset"`
}

type Rabbitmq struct {
	Brokers           string `mapstructure:"brokers"`
	ConsumerGroup     string `mapstructure:"consumer_group"`
	TopicAds          string `mapstructure:"topic_ads"`
	TopicActionStates string `mapstructure:"topic_action_states"`
	InitOffset        string `mapstructure:"init_offset"`
}

type Redis struct {
	URI               string `mapstructure:"uri"`
	ExpireTimeMinutes int    `mapstructure:"expire_time_minutes"`
}

var configs Config

func ServerConfig() Server {
	return configs.Server
}

func ServiceConfig() Service {
	return configs.Service
}

func MongodbConfig() Mongodb {
	return configs.Mongodb
}

func ElasticsearchConfig() Elasticsearch {
	return configs.Elasticsearch
}

func PostgresqlConfig() Postgresql {
	return configs.Postgresql
}

func KafkaConfig() Kafka {
	return configs.Kafka
}

func RabbitmqConfig() Rabbitmq {
	return configs.Rabbitmq
}

func RedisConfig() Redis {
	return configs.Redis
}

const (
	LastOffset = "LastOffset"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("error while reading config file, err: ", err)
	}
	err := viper.Unmarshal(&configs)
	if err != nil {
		fmt.Println("error while Unmarshal config, err: ", err)
	}
}
