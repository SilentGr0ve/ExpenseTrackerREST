package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SilentGr0ve/ExpenseTrackerREST/internal/core/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
	file *os.File
}

type loggerKey struct{}

var (
	key = loggerKey{}
)

func NewLogger(config config.LoggerConfig) (*Logger, error) {
	atomicLevel := zap.NewAtomicLevel()
	if err := atomicLevel.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, fmt.Errorf("unmarshall logger level: %w", err)
	}

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("mkdir log folder: %w", err)
	}
	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(config.Folder, fmt.Sprintf("%s.log", timestamp))

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("open logfile: %w", err)
	}

	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")
	consoleEncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)

	fileEncoderConfig := zap.NewDevelopmentEncoderConfig()
	fileEncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000")

	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), atomicLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), atomicLevel),
	)

	logger := zap.New(core, zap.AddCaller())

	return &Logger{
		Logger: logger,
		file:   logFile,
	}, nil
}

func FromContext(ctx context.Context) *Logger {
	logger, ok := ctx.Value(key).(*Logger)
	if !ok {
		panic("no logger in context")
	}

	return logger
}

func ToContext(ctx context.Context, log *Logger) context.Context {
	return context.WithValue(
		ctx,
		key,
		log,
	)
}

func (l *Logger) Close() error {
	if err := l.Logger.Sync(); err != nil {
		return fmt.Errorf("sync logger: %w", err)
	}
	return l.file.Close()
}

func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{
		Logger: l.Logger.With(fields...),
		file:   l.file,
	}
}
