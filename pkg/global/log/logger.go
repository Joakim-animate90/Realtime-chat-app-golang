package log

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

var (
	Logger  *zap.Logger
	String  = zap.String
	Any     = zap.Any
	Int     = zap.Int
	Float32 = zap.Float32
)

func InitLogger(logpath string, loglevel string) {

	hook := lumberjack.Logger{
		Filename:   logpath, // Log file path, default os.tempdir ()
		MaxSize:    100,     // Save 100m per log file, default 100M by default
		MaxBackups: 30,
		MaxAge:     7,
		Compress:   true,
	}
	write := zapcore.AddSync(&hook)
	// Set the log level
	// Debug can print Info Debug Warn
	// Info level can print warn info
	// Warn can only print warn
	// debug-> info-> warn-> error
	var level zapcore.Level
	switch loglevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	case "warn":
		level = zap.WarnLevel
	default:
		level = zap.InfoLevel
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	// Set the log level
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)

	var writes = []zapcore.WriteSyncer{write}
	// If it is the development environment, it also outputs it on the console
	if level == zap.DebugLevel {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		// zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...), // Print to the console and file
		// write,
		level,
	)
	// Open the development mode, stack tracking
	caller := zap.AddCaller()
	// Open the file and line number
	development := zap.Development()
	// Set the initialization field, such as: add a server name
	filed := zap.Fields(zap.String("application", "chat-room"))
	// Construction log
	Logger = zap.New(core, caller, development, filed)
	Logger.Info("Logger init success")
}
