package log

import (
	config "github.com/ct-core-standard/config"
	"go.uber.org/zap/zapcore"
	"time"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	logLevel := config.ServerConfig().LogLevel
	var configLogger zap.Config
	configLogger = zap.NewProductionConfig()
	if logLevel == "debug" {
		configLogger = zap.NewDevelopmentConfig()
	}
	configLogger.EncoderConfig.EncodeTime = syslogTimeEncoder
	configLogger.Level.UnmarshalText([]byte(logLevel))
	log, _ := configLogger.Build()
	logger = log.Named("service").Sugar()
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func GetLogger() *zap.SugaredLogger {
	return logger
}
