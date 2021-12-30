//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/zlog-fun/ysword"
	"ysword_layout/bootstrap"
	"ysword_layout/internal/biz"
	"ysword_layout/internal/repo"
	"ysword_layout/internal/router"
	"ysword_layout/internal/service"
)

func InitializeApp(bcf bootstrap.AppConfig) (*ysword.App, func(), error) {
	panic(wire.Build(
		repo.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		router.ProviderSet,
	newApp,
	))
}
