package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/google/wire"
)

type MysqlConf struct {
	Host   string `yaml:"host"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"password"`
	DbName string `yaml:"database"`
	Port   int    `yaml:"port"`
}

var ProviderMysqlConf = wire.NewSet(GetMysqlConf)

func GetMysqlConf() (*MysqlConf, error) {

	basePath, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "get pwd failed.")
	}

	fileName := filepath.Join(basePath, "configs", "db.yaml")
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "open db.yaml failed.")
	}

	m := &MysqlConf{}
	err = yaml.Unmarshal(yamlFile, m)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshal db.yaml failed.")
	}
	return m, nil
}
