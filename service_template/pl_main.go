package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	adaptix "github.com/Adaptix-Framework/axc2"
)

type PluginService struct{}

var (
	Ts        Teamserver
	ModuleDir string
	Config    ServiceConfig
)

func InitPlugin(ts any, moduleDir string, serviceConfig string) adaptix.PluginService {
	Ts = ts.(Teamserver)
	ModuleDir = moduleDir

	/// START CODE HERE

//	if err := loadConfig(serviceConfig); err != nil {
//		fmt.Printf("Config error: %v\n", err)
//		return &PluginService{}
//	}

	/// END CODE HERE

	return &PluginService{}
}

func (p *PluginService) Call(operator string, function string, args string) {
	/// START CODE HERE

	/// END CODE HERE
}

func loadConfig(serviceConfig string) error {
//	data, err := Ts.TsExtenderDataLoad("service_name", "key")
//	if err == nil && data != nil {
//		err = json.Unmarshal(data, &Config)
//		if err == nil {
//			return nil
//		}
//		fmt.Printf("Failed to load config: %v\n", err)
//		fmt.Printf("Use service configuration\n")
//	}
//	if serviceConfig == "" {
//		return fmt.Errorf("empty service config")
//	}
//	return yaml.Unmarshal([]byte(serviceConfig), &Config)
}

func saveConfig() error {
//	data, err := json.Marshal(&Config)
//	if err != nil {
//		return err
//	}
//	return Ts.TsExtenderDataSave("service_name", "key", data)
}