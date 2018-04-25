package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Config struct {
	Host     string `json:"host"`
	Port     string    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c Config) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", c.Host, c.Port, c.User, c.Database)
}

func NewConfig(filename string) (Config, error) {
	var cf Config
	confFile, err := os.Open("./" + filename)
	defer confFile.Close()
	
	if err != nil {
		return cf, err
	}
	jsonParser := json.NewDecoder(confFile)
	jsonParser.Decode(&cf)
	if err != nil {
		return cf, err
	}
	
	return cf, nil
}