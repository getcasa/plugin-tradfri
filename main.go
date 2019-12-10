package main

import (
	"encoding/json"
	"fmt"

	"github.com/getcasa/sdk"
)

// Config define the casa plugin configuration
var Config = sdk.Configuration{
	Name:        "Ikea Tradfri",
	Version:     "1.0.0",
	Author:      "casa",
	Description: "Control Ikea Tradfri ecosystem",
	Discover:    true,
	Devices: []sdk.Device{
		sdk.Device{
			Name:           "TRADFRI bulb E27 WS opal 1000lm",
			DefaultTrigger: "",
			DefaultAction:  "toggle",
			Triggers:       []sdk.Trigger{},
			Actions:        []string{"toggle"},
		},
	},
	Actions: []sdk.Action{
		sdk.Action{
			Name:   "toggle",
			Fields: []sdk.Field{},
		},
	},
}

func main() {
	OnStart([]byte("{}"))
}

type savedConfig struct {
	Identity string
	PSK      string
}

// ConfigPlugin store the saved config
var ConfigPlugin savedConfig

// Init plugin config
func Init() []byte {
	res, _ := json.Marshal([]savedConfig{})
	return res
}

// OnStart discover brdiges and create the global state
func OnStart(config []byte) {
	if err := json.Unmarshal(config, &ConfigPlugin); err != nil {
		panic(err)
	}

	Connect()
}

// Discover return array of all found devices
func Discover() []sdk.DiscoveredDevice {
	var discovered []sdk.DiscoveredDevice

	return discovered
}

// Params define actions parameters available
type Params struct {
	State bool `json:"state"`
}

// CallAction call functions from actions
func CallAction(physicalID string, name string, params []byte, config []byte) {
	if string(params) == "" {
		fmt.Println("Params must be provided")
	}

	// declare parameters
	var req Params

	// unmarshal parameters to use in actions
	err := json.Unmarshal(params, &req)
	if err != nil {
		fmt.Println(err)
	}

	// use name to call actions
	switch name {
	case "toggle":
		//
	default:
		return
	}
}

// OnStop close connection
func OnStop() {
	Conn.Close()
}
