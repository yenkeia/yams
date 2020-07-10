package game

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config 对应配置文件 yams.yaml
type Config struct {
	Assets string `yaml:"assets"`
	Mysql  struct {
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		DataDB    string `yaml:"dataDB"`
		AccountDB string `yaml:"accountDB"`
		Username  string `yaml:"username"`
		Password  string `yaml:"password"`
	}
}

// NewConfig ..
func NewConfig(path string) (conf *Config) {
	conf = new(Config)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
	return
}
