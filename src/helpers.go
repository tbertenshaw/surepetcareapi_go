package main

import (
	//   "fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

const configPath = "config.yml"

type Cfg struct {
	UrlDoorRoot      string `yaml:"urlDoorRoot"`
	BearerToken      string `yaml:"bearerToken"`
	UrlPetStatusRoot string `yaml:"urlPetStatusRoot"`
}

var AppConfig *Cfg

func (c *Cfg) GetConf() *Cfg {

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
