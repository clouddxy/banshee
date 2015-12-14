// Copyright 2015 Eleme Inc. All rights reserved.

package config

import (
	"encoding/json"
	"io/ioutil"
)

// Defaults
const (
	DefaultInterval     int     = 10
	DefaultNumGrids     int     = 288
	DefaultGridLen      int     = 300
	DefaultWeightFactor float64 = 0.05
	DefaultStartSize    int     = 18
)

type Config struct {
	Interval int            `json:"interval"`
	Period   [2]int         `json:"period"`
	Storage  ConfigStorage  `json:"storage"`
	Detector ConfigDetector `json:"detector"`
	Webapp   ConfigWebapp   `json:"webapp"`
	Alerter  ConfigAlerter  `json:"alerter"`
}

type ConfigStorage struct {
	Path string `json:"path"`
}

type ConfigDetector struct {
	Port      int      `json:"port"`
	Factor    float64  `json:"factor"`
	BlackList []string `json:"blackList"`
	StartSize int      `json:"startSize"`
}

type ConfigWebapp struct {
	Port int       `json:"port"`
	Auth [2]string `json:"auth"`
}

type ConfigAlerter struct {
	Command string `json:"command"`
	Workers int    `json:"workers"`
}

// New creates a Config with default values.
func New() *Config {
	config := new(Config)
	config.Interval = DefaultInterval
	config.Period = [2]int{DefaultNumGrids, DefaultGridLen}
	config.Storage.Path = "storage/"
	config.Detector.Port = 2015
	config.Detector.Factor = DefaultWeightFactor
	config.Detector.BlackList = []string{}
	config.Detector.StartSize = DefaultStartSize
	config.Webapp.Port = 2016
	config.Webapp.Auth = [2]string{"admin", "admin"}
	config.Alerter.Command = ""
	config.Alerter.Workers = 4
	return config
}

// Update config with json file.
func (config *Config) UpdateWithJsonFile(fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, config)
	if err != nil {
		return err
	}
	return err
}
