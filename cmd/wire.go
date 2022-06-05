//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"goHomework4/internal/biz"
	"goHomework4/internal/config"
	"goHomework4/internal/data"
)

func InitApp() (*App, error) {
	panic(wire.Build(config.ProviderMysqlConf, data.Provider, biz.NewApp))
}
