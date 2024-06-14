package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var SugaredLogger *zap.SugaredLogger

func InitLogger() {
	encoder := getEncoder()
	fileCore := zapcore.NewCore(encoder, getLogWriter(), getLevel())
	consoleCore := zapcore.NewCore(encoder, getConsoleWriter(), getLevel())
	core := zapcore.NewTee(fileCore, consoleCore)
	// dev zap.New(core, zap.AddCaller())
	logger := zap.New(core, zap.AddCaller())
	SugaredLogger = logger.Sugar()
}

// 根据生产环境记录日志级别
func getLevel() zapcore.Level {
	return zapcore.InfoLevel
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006/1/2 15:04:05")
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getConsoleWriter() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./logs/test.log", // 日志文件的位置
		MaxSize:    1,                 // 切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,                 // 日志保留数量
		MaxAge:     30,                // 日志保留期限
		Compress:   false,             // 日志文件压缩
		LocalTime:  true,              // false 使用 世界标准时间 进行切割. true 使用 计算机本地 时间
	}
	return zapcore.AddSync(lumberJackLogger)
}
