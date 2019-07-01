package conf

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type database struct {
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Port         int    `yaml:"port"`
	Host         string `yaml:"host"`
	DatabaseName string `yaml:"database"`
	Charset      string `yaml:"charset"`
}

type storage struct {
	Path string `yaml:"path"`
}

type mail struct {
	Address string `yaml:"address"`
	Token   string `yaml:"token"`
}

type AppConfig struct {
	DatabaseConfig database `yaml:"Database"`
	StorageConfig  storage  `yaml:"Storage"`
	MailConifg     mail     `yaml:"Mail"`
}

func InitConfig() *AppConfig {
	data, err := ioutil.ReadFile("conf/config.yml")
	if err != nil {
		log.WithError(err).Fatal("Read config file failed")
	}
	config := AppConfig{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.WithError(err).Fatal("Unmarshal config data failed")
	}
	return &config
}
