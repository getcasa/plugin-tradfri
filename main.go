package main

import (
	"fmt"
	"net"

	"github.com/dustin/go-coap"
	"github.com/pion/dtls"
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
	addr := &net.UDPAddr{IP: net.ParseIP(""), Port: 5684}
	config := &dtls.Config{
		PSK: func(_ []byte) ([]byte, error) {
			return []byte(""), nil
		},
		PSKIdentityHint: []byte(""), // For Tradfri Gateway the IdentityHint MUST be Client_identity
		CipherSuites:    []dtls.CipherSuiteID{dtls.TLS_PSK_WITH_AES_128_CCM_8},
	}

	dtlsConn, err := dtls.Dial("udp", addr, config)
	if err != nil {
		panic(err)
	}
	defer dtlsConn.Close()

	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 1,
	} // This is a CoAP Ping, empty confirmable message

	req.SetPathString("/15001")

	data, err := req.MarshalBinary()
	if err != nil {
		panic(err)
	}

	_, err = dtlsConn.Write(data)
	if err != nil {
		panic(err)
	}

	resp := make([]byte, 2048)
	dtlsConn.Read(resp)
	msg, err := coap.ParseMessage(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg.MessageID)
	fmt.Println(msg.Type)
	fmt.Println(msg.Code)
	fmt.Println(msg.Token)
	fmt.Println(string(msg.Payload))
}

type savedConfig struct {
	Identity string
	PSK string
}

// States is the global state of plugin
var States []State
var client http.Client
var configPlugin []savedConfig

// Init plugin config
func Init() []byte {
	res, _ := json.Marshal([]savedConfig{})
	return res
}

// Discover return array of all found devices
func Discover() []sdk.DiscoveredDevice {
	var discovered []sdk.DiscoveredDevice

	return discovered
}

// Params define actions parameters available
type Params struct {
	State  bool `json:"state"`
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
}
