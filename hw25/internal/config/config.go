package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host 				string `yaml:"Host"`
	Port 				string `yaml:"Port"`
	LogLevel 			string `yaml:"LogLevel"`
	FileName 			string `yaml:"FileName"`
	GrpcServer 			string `yaml:"GrpcServer"`
	DBServer	       	string `yaml:"DBServer"`
	DBUser           	string `yaml:"DBUser"`
	DBPass           	string `yaml:"DBPass"`
	DBDatabase       	string `yaml:"DBDatabase"`
	DBTimeoutConnect 	int    `yaml:"DBTimeoutConnect"`
	DBTimeoutExecute 	int    `yaml:"DBTimeoutExecute"`
	DBType 				string  `yaml:"DBType"`
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
