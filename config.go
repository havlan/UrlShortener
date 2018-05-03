package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
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

	if confFile != nil && err != nil {
		fmt.Printf("Creating config.")
		defer confFile.Close()

		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("%s", "File does not exist.")
				return cf, err
			} else {
				return cf, err
			}

		}
		jsonParser := json.NewDecoder(confFile)
		jsonParser.Decode(&cf)
		if err != nil {
			return cf, err
		}

		return cf, nil
	} else {
		return cf, errors.New("Could not find config.json")
	}
}
