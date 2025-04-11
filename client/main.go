package main

import (
	"net"

	"github.com/charmbracelet/log"
)

func main() {
	sendMessage("test", 1, firstType)
	sendMessage("wrong version", 2, firstType)
	sendMessage("wrong type", 1, thirdType)
}

func sendMessage(s string, version int8, messageType int8) {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatal("Connection error :", err)
		return
	}
	defer conn.Close()

	msg := message{messageVersion: version, messageType: messageType, messageContent: s}
	m := msg.toBytes()
	_, err = conn.Write(m)
	if err != nil {
		log.Fatal("An error has occured sending the message :", err)
		return
	}
}

type message struct {
	messageVersion int8
	messageType    int8
	messageContent string
}

func (m message) toBytes() []byte {
	return append([]byte{byte(m.messageVersion), byte(m.messageType), byte(len(m.messageContent))}, []byte(m.messageContent)...)
}

const firstType = int8(1)
const secondType = int8(2)
const thirdType = int8(3)
