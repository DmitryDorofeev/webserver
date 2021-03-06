package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Port        int    `json:"port"`
	Root        string `json:"root"`
	EnableLog   bool   `json:"enableLog"`
	DefaultFile string `json:"defaultFile"`
}

var config Config

func init() {
	config = readConfig()
}

func Get() Config {
	return config
}

func readConfig() Config {
	str, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("error in file")
	}
	res := &Config{}
	json.Unmarshal([]byte(str), &res)

	return *res
}
