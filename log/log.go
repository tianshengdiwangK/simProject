package setup

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

/*
   setJSONEncoder 设置logger编码
*/
func setJSONEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   //转换编码的时间戳
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //编码级别调整为大写的级别输出
	return zapcore.NewConsoleEncoder(encoderConfig)
}

/*
   setLoggerWrite 设置logger写入文件
*/
func setLoggerWrite() zapcore.WriteSyncer {
	l := &lumberjack.Logger{
		Filename:   "./logs/test.log",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     30,
		Compress:   true,
		LocalTime:  true,
	}
	return zapcore.AddSync(l)
}

/*
   InitLogger 初始化 logger
*/
func InitLogger() {
	core := zapcore.NewCore(setJSONEncoder(), setLoggerWrite(), zap.InfoLevel)
	logger = zap.New(core, zap.AddCaller()).Sugar()
}
func CwLog() *zap.SugaredLogger {
	return logger
}
