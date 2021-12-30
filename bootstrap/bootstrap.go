package bootstrap

import (
	"github.com/zlog-fun/ysword/config"
	"github.com/zlog-fun/ysword/graceful"
	"github.com/zlog-fun/ysword/logger"
	"github.com/zlog-fun/ysword/server"
)

type LoggerConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

// AppConfig global
type AppConfig struct {
	AppName  string                  `yaml:"appname"`
	Debug    string                  `yaml:"debug"`
	Pid      string                  `yaml:"pid"`
	Server   server.HttpConfig       `yaml:"server"`
	Logger   LoggerConfig            `yaml:"logger"`
	Graceful graceful.GracefulConfig `yaml:"graceful"`
}

type Bootstrap struct{}

func Setup(name, conf, env string) (AppConfig, error) {
	// load config
	cfg := config.New(config.Name(name), config.Path(conf), config.Env(env))
	if err := cfg.Load(); err != nil {
		panic(err)
	}
	var bcf AppConfig
	if err := cfg.Scan(&bcf); err != nil {
		panic("load config failure!")
	}
	// 初始化日志类
	l := logger.New(
		logger.Filename(bcf.Logger.Filename),
		logger.Level(bcf.Logger.Level),
		logger.MaxAge(bcf.Logger.MaxAge),
		logger.MaxSize(bcf.Logger.MaxSize),
		logger.Maxbackups(bcf.Logger.MaxBackups),
	)
	err := l.InitLogger()
	if err != nil {
		return bcf, err
	}

	return bcf, nil
}
