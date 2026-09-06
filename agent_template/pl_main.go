package main

import (
	"encoding/json"
	"errors"
	"math/rand/v2"
	"time"

	adaptix "github.com/Adaptix-Framework/axc2/v2"
)

type PluginAgent struct{}

var (
	Ts             adaptix.Teamserver
	ModuleDir      string
	AgentWatermark string
)

func InitPlugin(ts any, moduleDir string, watermark string) adaptix.PluginAgent {
	ModuleDir = moduleDir
	AgentWatermark = watermark
	Ts = ts.(adaptix.Teamserver)
	return &PluginAgent{}
}

func (p *PluginAgent) AgentRestore(agentData adaptix.AgentData) adaptix.AgentFunctions {
	return adaptix.AgentFunctions{
		CreateCommand: CreateCommand,
		ProcessData:   ProcessData,
		Encrypt:       Encrypt,
		Decrypt:       Decrypt,
		PackTasks:     PackTasks,
		PivotPackData: PivotPackData,
		TunnelCB: adaptix.TunnelCallbacks{
			ConnectTCP: TunnelMessageConnectTCP,
			ConnectUDP: TunnelMessageConnectUDP,
			WriteTCP:   TunnelMessageWriteTCP,
			WriteUDP:   TunnelMessageWriteUDP,
			Pause:      TunnelMessagePause,
			Resume:     TunnelMessageResume,
			Close:      TunnelMessageClose,
			Reverse:    TunnelMessageReverse,
			BindTCP:    TunnelMessageBindTCP,
		},
		TerminalCB: adaptix.TerminalCallbacks{
			Start: TerminalMessageStart,
			Write: TerminalMessageWrite,
			Close: TerminalMessageClose,
		},
	}
}

/// TUNNEL

func TunnelMessageConnectTCP(channelId int64, tunnelType int, addressType int, address string, port int) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageConnectUDP(channelId int64, tunnelType int, addressType int, address string, port int) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageWriteTCP(channelId int64, data []byte) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageWriteUDP(channelId int64, data []byte) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessagePause(channelId int64) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageResume(channelId int64) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageClose(channelId int64) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageReverse(tunnelId int64, port int) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TunnelMessageBindTCP(channelId int64, addressType int, address string, port int) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

/// TERMINAL

func TerminalMessageStart(terminalId int64, program string, sizeH int, sizeW int, oemCP int) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TerminalMessageWrite(terminalId int64, oemCP int, data []byte) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

func TerminalMessageClose(terminalId int64) adaptix.TaskData {
	var packData []byte
	/// START CODE HERE

	/// END CODE HERE
	return adaptix.MakeProxyTask(packData, 0)
}

/// BUILD

func (p *PluginAgent) GenerateProfiles(profile adaptix.BuildProfile) ([][]byte, error) {
	var agentProfiles [][]byte

	for _, transportProfile := range profile.ListenerProfiles {

		var listenerMap map[string]any
		if err := json.Unmarshal(transportProfile.Profile, &listenerMap); err != nil {
			return nil, err
		}

		/// START CODE HERE
		_ = listenerMap
		_ = AgentWatermark
		agentProfiles = append(agentProfiles, nil)
		/// END CODE HERE
	}
	return agentProfiles, nil
}

func (p *PluginAgent) BuildPayload(profile adaptix.BuildProfile, agentProfiles [][]byte) ([]byte, string, error) {
	var (
		Filename string
		Payload  []byte
	)

	/// START CODE HERE
	// profile.AgentConfig is JSON from GenerateUI container.
	_ = profile
	_ = agentProfiles
	/// END CODE HERE

	return Payload, Filename, nil
}

func (p *PluginAgent) CreateAgent(beat []byte) (adaptix.AgentData, adaptix.AgentFunctions, error) {
	var agentData adaptix.AgentData

	/// START CODE HERE
	_ = beat
	/// END CODE HERE

	return agentData, p.AgentRestore(agentData), nil
}

/// CRYPTO / TASK

func Encrypt(data []byte, key []byte) ([]byte, error) {
	/// START CODE HERE
	return data, nil
	/// END CODE HERE
}

func Decrypt(data []byte, key []byte) ([]byte, error) {
	/// START CODE HERE
	return data, nil
	/// END CODE HERE
}

func PackTasks(agentData adaptix.AgentData, tasks []adaptix.TaskData) ([]byte, error) {
	var packData []byte

	/// START CODE HERE
	// Encode tasks for the agent wire format. TaskId is int64.
	_ = agentData
	_ = tasks
	/// END CODE HERE

	return packData, nil
}

func PivotPackData(pivotId string, data []byte) (adaptix.TaskData, error) {
	var (
		packData []byte
		err      error
	)

	/// START CODE HERE
	_ = pivotId
	_ = data
	/// END CODE HERE

	return adaptix.TaskData{
		TaskId: int64(rand.Uint32()),
		Type:   adaptix.TASK_TYPE_PROXY_DATA,
		Data:   packData,
		Sync:   false,
	}, err
}

func CreateCommand(agentData adaptix.AgentData, args map[string]any) (adaptix.TaskData, adaptix.ConsoleMessageData, error) {
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
		Type: adaptix.TASK_TYPE_TASK,
		Sync: true,
	}

	messageData = adaptix.ConsoleMessageData{
		Status: adaptix.MESSAGE_INFO,
		Text:   "",
	}
	messageData.Message, _ = args["message"].(string)

	/// START CODE HERE
	_ = agentData
	_ = command
	_ = subcommand
	/// END CODE HERE

	return taskData, messageData, err
}

func ProcessData(agentData adaptix.AgentData, decryptedData []byte) error {
	var outTasks []adaptix.TaskData

	taskData := adaptix.TaskData{
		Type:        adaptix.TASK_TYPE_TASK,
		AgentId:     agentData.Id,
		FinishDate:  time.Now().Unix(),
		MessageType: adaptix.MESSAGE_SUCCESS,
		Completed:   true,
		Sync:        true,
	}

	/// START CODE HERE
	_ = decryptedData
	_ = taskData
	/// END CODE HERE

	for _, task := range outTasks {
		Ts.TsTaskUpdate(agentData.Id, task)
	}

	return nil
}

func (p *PluginAgent) Call(operator string, agentId int64, function string, args string) {
	_ = operator
	_ = agentId
	_ = function
	_ = args
}
