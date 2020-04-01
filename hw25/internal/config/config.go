package config

import (
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host 				string `yaml:"Host"`
	Port 				string `yaml:"Port"`
	LogLevel 			string `yaml:"LogLevel"`
	FileName 			string `yaml:"FileName"`
	GrpcServer 			string `yaml:"GrpcServer"`
	DB 					storage.StorageConfig	`yaml:"db"`
	Rmq       			rabbitmq.Config	`yaml:"rmq"`
}

func ReadFromFile(file string) (Config, error) {
	c := Config{}
	r, e := ioutil.ReadFile(file)
	if e != nil {
		return c, e
	}

	e = yaml.Unmarshal(r, &c)
	if e != nil {
		return c, e
	}

	return c, nil
}
