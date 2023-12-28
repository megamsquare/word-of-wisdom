package services

import (
	"fmt"
	"io"
	"net"
)

func sendServerMessage(msg Message, conn net.Conn) error {
	messageString := fmt.Sprintf("%s\n", msg.ConvertToString())
	_, err := conn.Write([]byte(messageString))
	return err
}

func sendClientMessage(msg Message, conn io.Writer) error {
	messageString := fmt.Sprintf("%s\n", msg.ConvertToString())
	_, err := conn.Write([]byte(messageString))
	return err
}