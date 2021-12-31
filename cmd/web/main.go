package main

import (
	"flag"
	"ysword_layout/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/zlog-fun/ysword"
)

// go build ./"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	conf string
	// 当前项目运行环境
	env string
)

func init() {
	flag.StringVar(&Name, "name", "blog", "application name, eg: -name blog")
	flag.StringVar(&conf, "conf", "./configs", "config path, eg: -conf ../../config.yaml")
	flag.StringVar(&env, "env", "", "enviroment load different config, default is null string, eg: -env test")
}

// config bootstrap.AppConfig
func newApp(c bootstrap.AppConfig, r func(engine *gin.Engine)) *ysword.App {
	return ysword.New(
		ysword.Name(Name),
		ysword.Version(Version),
		ysword.Handle(r),
		ysword.HttpConfig(c.Server),
	)
}

func main() {
	flag.Parse()
	bcf, err := bootstrap.Setup(Name, conf, env)
	if err != nil {
		panic(err.Error())
	}
	app, cleanup, err := InitializeApp(bcf)
	if err != nil {
		panic("initialize app failure")
	}
	defer cleanup()
	app.Run()
}
