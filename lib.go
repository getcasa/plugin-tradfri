package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/ItsJimi/go-arp"
	"github.com/dustin/go-coap"
	"github.com/pion/dtls"
)

// Conn is the udp connection
var Conn *dtls.Conn

// Connect establish connection with Ikea Tradfri gateway
func Connect() {
	entries, err := arp.GetEntries()
	if err != nil {
		return
	}

	var ip string
	for _, entry := range entries {
		if entry.HWAddress == strings.ReplaceAll(ConfigPlugin.Identity, "-", ":") {
			ip = entry.IPAddress
		}
	}

	addr := &net.UDPAddr{IP: net.ParseIP(ip), Port: 5684}
	config := &dtls.Config{
		PSK: func(_ []byte) ([]byte, error) {
			return []byte(ConfigPlugin.PSK), nil
		},
		PSKIdentityHint: []byte(ConfigPlugin.Identity), // For Tradfri Gateway the IdentityHint MUST be Client_identity
		CipherSuites:    []dtls.CipherSuiteID{dtls.TLS_PSK_WITH_AES_128_CCM_8},
	}

	Conn, err = dtls.Dial("udp", addr, config)
	if err != nil {
		panic(err)
	}
}

// GetDevices list all devices connected to Ikea Tradfri gateway
func GetDevices() {
	req := coap.Message{
		Type:      coap.Confirmable,
		Code:      coap.GET,
		MessageID: 1,
	}

	req.SetPathString("/15001")

	data, err := req.MarshalBinary()
	if err != nil {
		panic(err)
	}

	_, err = Conn.Write(data)
	if err != nil {
		panic(err)
	}

	resp := make([]byte, 2048)
	Conn.Read(resp)
	msg, err := coap.ParseMessage(resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(msg.Payload))
}
