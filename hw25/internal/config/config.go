package config

import (
	grpcserver "github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/grpc"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/logger"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/rabbitmq"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/scheduler"
	"github.com/a1ekaeyVorobyev/otus_go_hw/hw25/internal/storage"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host		string `yaml:"Host"`
	Port		string `yaml:"Port"`
	Log			logger.Config `yaml:"log"`
	Grps		grpcserver.Config `yaml:"grps"`
	DB 			storage.Config	`yaml:"db"`
	Rmq 		rabbitmq.Config	`yaml:"rmq"`
	Sheduler 	scheduler.Config `yaml:"scheduler"`
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
