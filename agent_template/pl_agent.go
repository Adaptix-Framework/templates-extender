package main

import (
	"encoding/json"
	"errors"

	"github.com/Adaptix-Framework/axc2"
)

type GenerateConfig struct {
	// CONFIG PARAMS HERE
}

func AgentGenerateProfile(agentConfig string, listenerWM string, listenerMap map[string]any) ([]byte, error) {
	var (
		generateConfig GenerateConfig
		err            error
		params         []interface{}
		profileString  string
	)

	err = json.Unmarshal([]byte(agentConfig), &generateConfig)
	if err != nil {
		return nil, err
	}

	/// START CODE HERE

	/// END CODE HERE

	return []byte(profileString), nil
}

func AgentGenerateBuild(agentConfig string, agentProfile []byte, listenerMap map[string]any) ([]byte, string, error) {
	var (
		generateConfig GenerateConfig
		Filename       string
		buildContent   []byte
	)

	err := json.Unmarshal([]byte(agentConfig), &generateConfig)
	if err != nil {
		return nil, "", err
	}

	/// START CODE HERE

	/// END CODE HERE

	return buildContent, Filename, nil
}

func CreateAgent(initialData []byte) (adaptix.AgentData, error) {
	var agentData adaptix.AgentData

	/// START CODE HERE

	/// END CODE

	return agentData, nil
}

func AgentEncryptData(data []byte, key []byte) ([]byte, error) {
	/// START CODE
	return data, nil
	/// END CODE
}

func AgentDecryptData(data []byte, key []byte) ([]byte, error) {
	/// START CODE
	return data, nil
	/// END CODE
}

/// TASKS

func PackTasks(agentData adaptix.AgentData, tasksArray []adaptix.TaskData) ([]byte, error) {
	var packData []byte

	/// START CODE HERE

	/// END CODE

	return packData, nil
}

func PackPivotTasks(pivotId string, data []byte) ([]byte, error) {
	/// START CODE HERE
	return data, nil
	/// END CODE HERE
}

func CreateTask(ts Teamserver, agent adaptix.AgentData, args map[string]any) (adaptix.TaskData, adaptix.ConsoleMessageData, error) {
	var (
		taskData    adaptix.TaskData
		messageData adaptix.ConsoleMessageData
		err         error
	)

	command, ok := args["command"].(string)
	if !ok {
		return taskData, messageData, errors.New("'command' must be set")
	}
	subcommand, _ := args["subcommand"].(string)

	taskData = adaptix.TaskData{
		Type: TYPE_TASK,
		Sync: true,
	}

	messageData = adaptix.ConsoleMessageData{
		Status: MESSAGE_INFO,
		Text:   "",
	}
	messageData.Message, _ = args["message"].(string)

	/// START CODE HERE

	/// END CODE

	return taskData, messageData, err
}

func ProcessTasksResult(ts Teamserver, agentData adaptix.AgentData, taskData adaptix.TaskData, packedData []byte) []adaptix.TaskData {
	var outTasks []adaptix.TaskData

	/// START CODE

	/// END CODE

	return outTasks
}

/// TUNNELS

func TunnelCreateTCP(channelId int, tunnelType int, addressType int, address string, port int) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Tunnel not supported")
	/// END CODE HERE
}

func TunnelCreateUDP(channelId int, tunnelType int, addressType int, address string, port int) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Tunnel not supported")
	/// END CODE HERE
}

func TunnelWriteTCP(channelId int, data []byte) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Tunnel not supported")
	/// END CODE HERE
}

func TunnelWriteUDP(channelId int, data []byte) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Tunnel not supported")
	/// END CODE HERE
}

func TunnelClose(channelId int) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Tunnel not supported")
	/// END CODE HERE
}

func TunnelReverse(tunnelId int, port int) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Tunnel not supported")
	/// END CODE HERE
}

/// TERMINAL

func TerminalStart(terminalId int, program string, sizeH int, sizeW int, oemCP int) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Remote Terminal not supported")
	/// END CODE HERE
}

func TerminalWrite(terminalId int, oemCP int, data []byte) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Remote Terminal not supported")
	/// END CODE HERE
}

func TerminalClose(terminalId int) ([]byte, error) {
	/// START CODE HERE
	return nil, errors.New("Function Remote Terminal not supported")
	/// END CODE HERE
}
