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

type AccessLog struct {
	Ip            string
	RequestId     string
	Method        string
	Url           string
	RequestBody   string
	ResponseCode  int
	ResponseBody  string
	ResponserTime string
}

func RequestLog(ac AccessLog, filed ...zapcore.Field) {
	c := ac.ResponseCode
	defaultLog := log.With(
		zap.String("Ip", ac.Ip),
		zap.String("RequestId", ac.RequestId),
		zap.String("Method", ac.Method),
		zap.String("Url", ac.Url),
		zap.String("RequestBody", ac.RequestBody),
		zap.Int("ResponseCode", ac.ResponseCode),
		zap.String("ResponseBody", ac.ResponseBody),
		zap.String("ResponserTime", ac.ResponserTime),
	)
	if c >= 20000 && c <= 39999 {
		defaultLog.Info("success", filed...)
	} else if c >= 40000 && c <= 49999 {
		defaultLog.Info("warning", filed...)
	} else {
		defaultLog.Error("error", filed...)
	}
}

func Info(message string, filed ...zapcore.Field) {
	log.Info(message, filed...)
	// log.With(zap.String("func", funcName)).Info(message, filed...)
}

func Debug(message string, filed ...zapcore.Field) {
	log.Debug(message, filed...)
}

func Error(message interface{}, filed ...zapcore.Field) {
	switch t := message.(type) {
	case error:
		log.Error(t.Error(), filed...)
	case string:
		log.Error(t, filed...)
	}
}
