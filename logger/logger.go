package logger

import (
	"bufio"
	"bytes"
	"context"
	"os"

	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrzap"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(app *newrelic.Application) *Logger {
	core := zapcore.NewCore(zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), zapcore.AddSync(zapcore.Lock(os.Stdout)), zapcore.DebugLevel)
	backgroundcore, err := nrzap.WrapBackgroundCore(core, app)
	if err != nil && err != nrzap.ErrNilApp {
		panic(err)
	}
	backgroundLogger := zap.New(backgroundcore)
	return &Logger{SugaredLogger: backgroundLogger.Sugar()}
}

func NewTestLogger() (*Logger, *bytes.Buffer, *bufio.Writer) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	logger := zap.NewExample() // Create a new test logger instance

	return &Logger{
		SugaredLogger: logger.Sugar(),
	}, &buffer, writer
}

// DebugContext uses fmt.Sprint to log a message.
func (logger *Logger) DebugContext(ctx context.Context, args ...interface{}) {
	// If you want to log context data, you can extract from ctx
	logger.Debug(args...)
}

// DebugfContext uses fmt.Sprintf to log a templated message.
func (logger *Logger) DebugfContext(ctx context.Context, template string, args ...interface{}) {
	logger.Debugf(template, args...)
}

// ErrorContext uses fmt.Sprint to log a message.
func (logger *Logger) ErrorContext(ctx context.Context, args ...interface{}) {
	logger.Error(args...)
}

// ErrorfContext uses fmt.Sprintf to log a templated message.
func (logger *Logger) ErrorfContext(ctx context.Context, template string, args ...interface{}) {
	logger.Errorf(template, args...)
}

// InfoContext uses fmt.Sprint to log a message.
func (logger *Logger) InfoContext(ctx context.Context, args ...interface{}) {
	logger.Info(args...)
}

// InfofContext uses fmt.Sprintf to log a templated message.
func (logger *Logger) InfofContext(ctx context.Context, template string, args ...interface{}) {
	logger.Infof(template, args...)
}

// WarnContext uses fmt.Sprint to log a message.
func (logger *Logger) WarnContext(ctx context.Context, args ...interface{}) {
	logger.Warn(args...)
}

// WarnfContext uses fmt.Sprintf to log a templated message.
func (logger *Logger) WarnfContext(ctx context.Context, template string, args ...interface{}) {
	logger.Warnf(template, args...)
}
