package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

var config = "config.yml"

type Config struct {
	InputFile      string
	OutputFile     string
	MasterSheet    string
	InboundSheet   string
	OutboundSheet  string
	InventorySheet string
}

func GetConfig() (Config, error) {

	byt, err := ioutil.ReadFile(config)
	if err != nil {

		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(byt, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}
