package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var SugaredLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, getLevel())
	logger := zap.New(core, zap.AddCaller())
	SugaredLogger = logger.Sugar()

}

func Fatal(error ...any) {
	SugaredLogger.Fatal(error)
}

func Error(error ...any) {
	SugaredLogger.Error(error)
}

func Info(error ...any) {
	SugaredLogger.Info(error)
}

// 根据生产环境记录日志级别
func getLevel() zapcore.Level {
	return zapcore.ErrorLevel
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("../logs/test.log"), // 日志文件的位置
		MaxSize:    1,                               // 切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,                               // 日志保留数量
		MaxAge:     30,                              // 日志保留期限
		Compress:   false,                           // 日志文件压缩
		LocalTime:  true,                            // false 使用 世界标准时间 进行切割. true 使用 计算机本地 时间
	}
	return zapcore.AddSync(lumberJackLogger)
}
