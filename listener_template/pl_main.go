package main

import (
	"bytes"
	"encoding/json"
	"strconv"

	adaptix "github.com/Adaptix-Framework/axc2/v2"
)

const (
	logSrc = "listener"
	logCtg = "_LISTENER_"
)

type PluginListener struct{}

var (
	ModuleDir       string
	ListenerDataDir string
	Ts              adaptix.Teamserver
)

func InitPlugin(ts any, moduleDir string, listenerDir string) adaptix.PluginListener {
	ModuleDir = moduleDir
	ListenerDataDir = listenerDir
	Ts = ts.(adaptix.Teamserver)
	return &PluginListener{}
}

/// CHANGE ME //////////////////////////////////////////////////

// TransportConfig is the JSON shape of ListenerUI container + persisted customData.
type TransportConfig struct {
	HostBind string `json:"host_bind"`
	PortBind int    `json:"port_bind"`
	Protocol string `json:"protocol"`
}

type Listener struct {
	name   string
	config TransportConfig
	active bool
}

//////////////////////////////////////////////////////////////

func (p *PluginListener) Create(name string, config string, customData []byte) (adaptix.ExtenderListener, adaptix.ListenerData, []byte, error) {
	var (
		listener     *Listener
		listenerData adaptix.ListenerData
		conf         TransportConfig
		customdData  []byte
		err          error
	)

	/// START CODE HERE

	if customData == nil {
		if err = json.Unmarshal([]byte(config), &conf); err != nil {
			return nil, listenerData, customdData, err
		}
		conf.Protocol = "_PROTOCOL_"
	} else {
		if err = json.Unmarshal(customData, &conf); err != nil {
			return nil, listenerData, customdData, err
		}
	}

	listener = &Listener{name: name, config: conf}
	listenerData = adaptix.ListenerData{
		Name:     name,
		Protocol: conf.Protocol,
		BindHost: conf.HostBind,
		BindPort: strconv.Itoa(conf.PortBind),
		Status:   "Stopped",
	}
	customdData, err = json.Marshal(conf)
	if err != nil {
		return nil, listenerData, customdData, err
	}

	/// END CODE HERE

	return listener, listenerData, customdData, nil
}

func (l *Listener) Start() error {
	/// START CODE HERE
	// return l.transport.Start(Ts)
	l.active = true
	return nil
	/// END CODE HERE
}

func (l *Listener) Edit(config string) (adaptix.ListenerData, []byte, error) {
	var (
		listenerData adaptix.ListenerData
		conf         TransportConfig
		customdData  []byte
		err          error
	)

	/// START CODE HERE

	if err = json.Unmarshal([]byte(config), &conf); err != nil {
		return listenerData, customdData, err
	}
	l.config.HostBind = conf.HostBind
	l.config.PortBind = conf.PortBind

	listenerData = adaptix.ListenerData{
		Name:     l.name,
		Protocol: l.config.Protocol,
		BindHost: l.config.HostBind,
		BindPort: strconv.Itoa(l.config.PortBind),
	}
	if l.active {
		listenerData.Status = "Listen"
	} else {
		listenerData.Status = "Closed"
	}
	customdData, err = json.Marshal(l.config)

	/// END CODE HERE

	return listenerData, customdData, err
}

func (l *Listener) Stop() error {
	/// START CODE HERE
	// return l.transport.Stop()
	l.active = false
	return nil
	/// END CODE HERE
}

func (l *Listener) GetProfile() ([]byte, error) {
	var buffer bytes.Buffer

	/// START CODE HERE
	// Blob consumed by the agent plugin GenerateProfiles.
	if err := json.NewEncoder(&buffer).Encode(l.config); err != nil {
		return nil, err
	}
	/// END CODE HERE

	return buffer.Bytes(), nil
}

func (l *Listener) InternalHandler(data []byte) (int64, error) {
	var agentId int64

	/// START CODE HERE
	// Internal (pivot) listeners: parse data and return agent id.
	_ = data
	/// END CODE HERE

	return agentId, nil
}

func (p *PluginListener) Call(operator string, listenerName string, function string, args string) {
	_ = operator
	_ = listenerName
	_ = function
	_ = args
}
