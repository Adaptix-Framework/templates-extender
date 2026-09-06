package main

import (
	adaptix "github.com/Adaptix-Framework/axc2/v2"
)

const (
	logSrc = "service"
	logCtg = "_SERVICE_"
)

type PluginService struct{}

var (
	Ts        adaptix.Teamserver
	ModuleDir string
	// Config    ServiceConfig
)

func InitPlugin(ts any, moduleDir string, serviceConfig string) adaptix.PluginService {
	Ts = ts.(adaptix.Teamserver)
	ModuleDir = moduleDir

	/// START CODE HERE

	// if err := loadConfig(serviceConfig); err != nil {
	// 	Ts.TsLogAdd(adaptix.LogStatusWarn, 0, logSrc, logCtg, "Config error: %v", err)
	// 	return &PluginService{}
	// }
	_ = serviceConfig

	/// END CODE HERE

	return &PluginService{}
}

func (p *PluginService) CallRPC(operator string, function string, args string) (string, error) {
	/// START CODE HERE

	// switch function {
	// case "get_config":
	// 	return `{"ok":true}`, nil
	// case "set_config":
	// 	return `{"ok":true}`, nil
	// }
	_ = operator
	_ = function
	_ = args
	return "", nil

	/// END CODE HERE
}

func (p *PluginService) Call(operator string, function string, args string) {
	res, err := p.CallRPC(operator, function, args)
	if err != nil {
		Ts.TsPluginServiceSendDataClient(operator, "_SERVICE_", `{"ok":false}`)
		return
	}
	if res != "" {
		Ts.TsPluginServiceSendDataClient(operator, "_SERVICE_", res)
	}
}

////////// Example config persistence helpers (uncomment and adapt):
//
// func loadConfig(serviceConfig string) error {
// 	data, err := Ts.TsExtenderDataLoad("_SERVICE_", "config")
// 	if err == nil && data != nil {
// 		if err = json.Unmarshal(data, &Config); err == nil {
// 			return nil
// 		}
// 	}
// 	if serviceConfig == "" {
// 		return fmt.Errorf("empty service config")
// 	}
// 	// return yaml.Unmarshal([]byte(serviceConfig), &Config)
// 	return nil
// }
//
// func saveConfig() error {
// 	data, err := json.Marshal(&Config)
// 	if err != nil {
// 		return err
// 	}
// 	return Ts.TsExtenderDataSave("_SERVICE_", "config", data)
// }
