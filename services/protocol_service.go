package services

import (
	"fmt"
	"strconv"
	"strings"
)

type Message struct {
	Header int
	Payload string
}

func ParseMsg(str string) (*Message, error) {
	str = strings.TrimSpace(str)

	splitStr := strings.Split(str, "|")
	if len(splitStr) < 1 || len(splitStr) > 2 {
		return nil, fmt.Errorf("invalid message")
	}

	strToNum, err := strconv.Atoi(splitStr[0])
	if err != nil {
		return nil, fmt.Errorf("error converting string to number")
	}

	msg := Message {
		Header : strToNum,
	}

	if len(splitStr) == 2 {
		msg.Payload = splitStr[1]
	}

	return &msg, nil
}

func (m *Message) ConvertToString() string {
	return fmt.Sprintf("%d|%s", m.Header, m.Payload)
}

