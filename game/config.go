package game

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type config Yaml

// Yaml 对应配置文件
type Yaml struct {
	Mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		DB       string `yaml:"db"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

func newConfig(path string) *config {
	conf = new(config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
	return conf
}
