package main

import (
	"sync"
	"log"
	"os"
	"encoding/json"
)

type config struct {
	Token    string `json:"token"`
	LogInfo  *log.Logger
	LogWarn  *log.Logger
	LogError *log.Logger
}

func (c *config) load() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	confFile, err := os.Open("./conf.json")
	if err != nil {
		log.Fatal(err)
	}
	dc := json.NewDecoder(confFile)
	if err := dc.Decode(&c); err != nil {
		log.Fatal("Read config file: ", err)
	}

	c.LogError = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	c.LogWarn = log.New(os.Stdout, "[WARNING] ", log.Ldate|log.Ltime)
	c.LogInfo = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

}

var _config *config
var _once sync.Once

func GetConfig() *config {
	_once.Do(func() {
		_config = new(config)
		_config.load()
	})
	return _config
}
