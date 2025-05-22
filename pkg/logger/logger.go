package logger

import (
	"context"
	"fmt"
	"ng_autocode_back_service/pkg/config"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logObject *zap.Logger

func Log(data ...context.Context) *zap.Logger {
	if data != nil {
		ctx := data[0]
		return logObject.With(zap.Any("requestID", ctx.Value(config.REQUESTID)), zap.Any("userID", ctx.Value(config.USERID)), zap.Any("ucc", ctx.Value(config.UCC)))
	} else {
		return logObject
	}
}

func LoggerInit(logFilePath string, level zapcore.Level) {
	// Initialize the logger with the specified log file path and level
	// Create a new logger instance with the specified log level and file path
	var err error = nil
	cfg := zap.NewProductionConfig()

	cfg.Level = zap.NewAtomicLevelAt(level)
	cfg.EncoderConfig.FunctionKey = "f"
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeTime = sylogTimeEncoder
	cfg.EncoderConfig.ConsoleSeparator = " "
	cfg.EncoderConfig.EncodeCaller = MyCaller
	cfg.Encoding = "console"
	cfg.DisableStacktrace = true

	if logFilePath != "" {
		var paths []string
		paths = append(paths, logFilePath)
		cfg.OutputPaths = paths
	}

	logObject, err := cfg.Build()
	if err != nil {
		fmt.Println("Error initializing logger:", err)
		os.Exit(1)
	} else if logObject == nil {
		logObject, err = zap.NewProduction()
		logObject.WithOptions(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.InfoLevel))
		if err != nil {
			fmt.Println("Error to create production logger:", err)
			os.Exit(1)
		}
		fmt.Println("Error to create custom production logger, Creating standard production logger", err)
		os.Exit(1)
	} else {
		// do nohing
	}
}

func GetZapLevel(level int) zapcore.Level {
	return zapcore.Level(level)
}

func sylogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("Jan 2 15:04:05.000"))
}

func customLevelEncode(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func MyCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(filepath.Base(caller.FullPath()))
}
