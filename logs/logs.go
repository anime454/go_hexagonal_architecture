package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.FunctionKey = "func"
	// config.EncoderConfig.MessageKey = "message"
	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fileds ...zapcore.Field) {
	log.Info(message, fileds...)
	// log.With(zap.String("func", funcName)).Info(message, fileds...)
}

func Debug(message string, fileds ...zapcore.Field) {
	log.Debug(message, fileds...)
}

func Error(message interface{}, fileds ...zapcore.Field) {
	switch t := message.(type) {
	case error:
		log.Error(t.Error(), fileds...)
	case string:
		log.Error(t, fileds...)
	}
}
