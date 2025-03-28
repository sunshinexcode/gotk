package vlog

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vreflect"
	"github.com/sunshinexcode/gotk/vtrace"
)

type (
	Log struct {
		Logger  *zap.SugaredLogger
		Options *Options
	}

	Options struct {
		Compress         bool
		File             string        // log file
		Level            zapcore.Level // log level
		MaxAge           int           // retention days
		MaxBackups       int           // MaxBackups is the maximum number of old log files to retain.
		MaxSize          int           // MaxSize is the maximum size in megabytes of the log file before it gets rotated. File size in MB
		OutputMode       OutputMode
		OutputFormatMode OutputFormatMode
	}

	OutputMode       uint
	OutputFormatMode uint
)

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	FatalLevel = zapcore.FatalLevel
)

const (
	OutputModeConsole OutputMode = iota
	OutputModeFile
	OutputModeConsoleAndFile
)

const (
	OutputFormatModeText OutputFormatMode = iota
	OutputFormatModeJson
)

var (
	log           *Log
	defaultOption = map[string]any{
		"Compress":         false,
		"File":             "/data/logs/app.log",
		"Level":            zap.DebugLevel,
		"MaxAge":           5,
		"MaxBackups":       10,
		"MaxSize":          500,
		"OutputMode":       OutputModeConsole,
		"OutputFormatMode": OutputFormatModeText,
	}
)

func init() {
	log, _ = New(nil)
}

func New(options map[string]any) (*Log, error) {
	log := &Log{Options: &Options{}}
	return log.SetConfig(options)
}

func GetLog() *Log {
	return log
}

func SetConfig(options map[string]any) (*Log, error) {
	return log.SetConfig(options)
}

func (log *Log) SetConfig(options map[string]any) (*Log, error) {
	if err := vreflect.SetAttrs(log.Options, vmap.Merge(defaultOption, options)); err != nil {
		return log, err
	}

	core := zapcore.NewCore(
		log.getEncoder(),
		zapcore.NewMultiWriteSyncer(log.getMultiWriteSyncer()...),
		zap.NewAtomicLevelAt(log.Options.Level),
	)

	log.Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)).Sugar()

	return log, nil
}

func (log *Log) getEncoder() (encoder zapcore.Encoder) {
	switch log.Options.OutputFormatMode {
	case OutputFormatModeText:
		encoder = zapcore.NewConsoleEncoder(log.getEncoderConfig())
	case OutputFormatModeJson:
		encoder = zapcore.NewJSONEncoder(log.getEncoderConfig())
	default:
		encoder = zapcore.NewConsoleEncoder(log.getEncoderConfig())
	}

	return
}

func (log *Log) getEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		CallerKey:        "caller",
		ConsoleSeparator: " - ",
		EncodeCaller:     zapcore.ShortCallerEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeLevel:      zapcore.CapitalLevelEncoder,
		EncodeTime:       zapcore.ISO8601TimeEncoder,
		LevelKey:         "level",
		LineEnding:       zapcore.DefaultLineEnding,
		MessageKey:       "msg",
		NameKey:          "logger",
		StacktraceKey:    "stacktrace",
		TimeKey:          "ts",
	}
}

func (log *Log) getMultiWriteSyncer() (multiWriteSyncer []zapcore.WriteSyncer) {
	switch log.Options.OutputMode {
	case OutputModeConsole:
		multiWriteSyncer = append(multiWriteSyncer, os.Stdout)
	case OutputModeFile:
		multiWriteSyncer = append(multiWriteSyncer, log.getWriteSyncer())
	case OutputModeConsoleAndFile:
		multiWriteSyncer = append(multiWriteSyncer, log.getWriteSyncer(), os.Stdout)
	default:
		multiWriteSyncer = append(multiWriteSyncer, os.Stdout)
	}

	return
}

func (log *Log) getWriteSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Compress:   log.Options.Compress,
		Filename:   log.Options.File,
		MaxAge:     log.Options.MaxAge,
		MaxBackups: log.Options.MaxBackups,
		MaxSize:    log.Options.MaxSize,
	}

	return zapcore.AddSync(lumberJackLogger)
}

func addTraceIdForArgs(ctx context.Context, args ...any) []any {
	args = append([]any{vtrace.GetTraceIdKey(), vtrace.GetTraceId(ctx)}, args...)
	return args
}

func Debug(msg string, args ...any) {
	log.Logger.Debugw(msg, args...)
}

func Debugf(msg string, args ...any) {
	log.Logger.Debugf(msg, args...)
}

func Debugc(ctx context.Context, msg string, args ...any) {
	log.Logger.Debugw(msg, addTraceIdForArgs(ctx, args...)...)
}

func Info(msg string, args ...any) {
	log.Logger.Infow(msg, args...)
}

func Infof(msg string, args ...any) {
	log.Logger.Infof(msg, args...)
}

func Infoc(ctx context.Context, msg string, args ...any) {
	log.Logger.Infow(msg, addTraceIdForArgs(ctx, args...)...)
}

func Warn(msg string, args ...any) {
	log.Logger.Warnw(msg, args...)
}

func Warnf(msg string, args ...any) {
	log.Logger.Warnf(msg, args...)
}

func Warnc(ctx context.Context, msg string, args ...any) {
	log.Logger.Warnw(msg, addTraceIdForArgs(ctx, args...)...)
}

func Error(msg string, args ...any) {
	log.Logger.Errorw(msg, args...)
}

func Errorf(msg string, args ...any) {
	log.Logger.Errorf(msg, args...)
}

func Errorc(ctx context.Context, msg string, args ...any) {
	log.Logger.Errorw(msg, addTraceIdForArgs(ctx, args...)...)
}

func Fatal(msg string, args ...any) {
	log.Logger.Fatalw(msg, args...)
}

func Fatalf(msg string, args ...any) {
	log.Logger.Fatalf(msg, args...)
}

func Fatalc(ctx context.Context, msg string, args ...any) {
	log.Logger.Fatalw(msg, addTraceIdForArgs(ctx, args...)...)
}
