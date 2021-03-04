package log

import (
	"github.com/google/wire"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// Options is log configuration struct
type Options struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Level      string
	Stdout     bool
}

/*
Development：bool 是否是开发环境。如果是开发模式，对DPanicLevel进行堆栈跟踪
DisableCaller：bool 禁止使用调用函数的文件名和行号来注释日志。默认进行注释日志
DisableStacktrace：bool 是否禁用堆栈跟踪捕获。默认对Warn级别以上和生产error级别以上的进行堆栈跟踪。
Encoding：编码类型，目前两种json 和 console【按照空格隔开】,常用json
EncoderConfig：生成格式的一些配置--TODO 后面我们详细看下EncoderConfig配置各个说明
OutputPaths：[]string 日志写入文件的地址
ErrorOutputPaths：[]string 将系统内的error记录到文件的地址
InitialFields：map[string]interface{} 加入一些初始的字段数据，比如项目名
当然了，如果想控制台输出，OutputPaths和ErrorOutputPaths不能配置为文件地址，而应该改为stdout。
 */

func NewOptions(v *viper.Viper) (*Options, error) {
	var (
		err error
		o   = new(Options)
	)
	if err = v.UnmarshalKey("log", o); err != nil {
		return nil, err
	}

	return o, err
}

// New for init zap log library
func New(o *Options) (*zap.Logger, error) {
	var (
		err    error
		level  = zap.NewAtomicLevel()
		logger *zap.Logger
	)

	err = level.UnmarshalText([]byte(o.Level))
	if err != nil {
		return nil, err
	}

	fw := zapcore.AddSync(&lumberjack.Logger{
		Filename:   o.Filename,
		MaxSize:    o.MaxSize, // megabytes
		MaxBackups: o.MaxBackups,
		MaxAge:     o.MaxAge, // days
	})

	cw := zapcore.Lock(os.Stdout)

	// file core 采用jsonEncoder
	cores := make([]zapcore.Core, 0, 2)
	je := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	cores = append(cores, zapcore.NewCore(je, fw, level))

	// stdout core 采用 ConsoleEncoder
	if o.Stdout {
		ce := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		cores = append(cores, zapcore.NewCore(ce, cw, level))
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core)

	zap.ReplaceGlobals(logger)

	return logger, err
}

var ProviderSet = wire.NewSet(New, NewOptions)

