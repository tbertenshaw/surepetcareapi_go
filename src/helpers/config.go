package helpers

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const configPath = "config.yml"

type Cfg struct {
	UrlDoorRoot        string `yaml:"urlDoorRoot"`
	BearerToken        string `yaml:"bearerToken"`
	UrlPetStatusRoot   string `yaml:"urlPetStatusRoot"`
	UrlPetLocationRoot string `yaml:"urlPetLocationRoot"`
}

var AppConfig *Cfg

func (c *Cfg) GetConf() *Cfg {

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// c is the internal variable 
//and this method is reachable of the struct definition and returns
//a struct with the config populated.
//
func (c *Cfg) ReadConfig() *Cfg{
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)

	if err != nil {
		fmt.Println(err)
	}
	return c
}
