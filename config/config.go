package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type MyConfig struct {
	Posqres string `yaml:"posqres" env-default:"postgres://grandpat:grandpat@localhost:5432/postgres"`
}

var instance *MyConfig
var once sync.Once

func GetConfig() *MyConfig {
	once.Do(func() {
		instance = &MyConfig{}
		if err := cleanenv.ReadConfig("../config/config.yml", instance); err != nil {
			log.Fatal(err)
		}
	})
	return instance
}
