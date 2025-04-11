package main

import (
	"io"
	"net"

	"github.com/charmbracelet/log"
)

var serverVersion byte = 1

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("Error Accepting connection: %v", err)
		}
		log.Info("Connection established " + conn.LocalAddr().String())
		go handleConnexion(conn)
	}
}

func handleConnexion(conn net.Conn) {
	defer conn.Close()

	for {
		header := make([]byte, 3)
		_, err := io.ReadFull(conn, header)
		if err != nil {
			if err == io.EOF {
				log.Info("Client disconnected.")
				return
			} else {
				log.Error("Error has occured when reading the header", err)
			}
			return
		}

		versionMessage := header[0]
		if versionMessage != serverVersion {
			log.Error("Version mismatch")
		}
		typeMessage := int8(header[1])
		lengthMessage := int8(header[2])

		content := make([]byte, lengthMessage)
		_, err = io.ReadFull(conn, content)
		if err != nil {
			log.Fatal("Errot reading the message:", err)
			return
		}

		switch typeMessage {
		case firstType:
			log.Info("Recived message of type 1: ", string(content))
			break
		case secondType:
			log.Info("Recieved message of type 2 : ", string(content))
			break
		default:
			log.Error("Unknown message type :", typeMessage)
		}

	}

}

type message struct {
	messageVersion int8 `default:"1"`
	messageType    int8
	messageContent string
}

const firstType = int8(1)
const secondType = int8(2)
