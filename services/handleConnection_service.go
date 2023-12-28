package services

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/megamsquare/word-of-wisdom/pkg/config"
)



func handleServerConnection(ctx context.Context, conn net.Conn) {
	// TODO: Handle the client request and send back a response

	fmt.Println("new client: ", conn.RemoteAddr())
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		req, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read connection err: ", err)
			return
		}

		msg, err := ProcessRequest(ctx, req, conn.RemoteAddr().String())
		if err != nil {
			fmt.Println("error processing request: ", err)
		}
		if msg != nil {
			err := sendServerMessage(*msg, conn)
			if err != nil {
				fmt.Println("error sending message: ", err)
			}
		}
	}
}

func handleClientConnection(ctx context.Context, readConn io.Reader, writeConn io.Writer) (string, error) {
	reader := bufio.NewReader(readConn)

	err := sendClientMessage(Message{
		Header: RequestChallenge,
	}, writeConn)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}

	msgStr, err := readConnMsg(reader)
	if err != nil {
		return "", fmt.Errorf("error reading message: %w", err)
	}

	msg, err := ParseMsg(msgStr)
	if err != nil {
		return "", fmt.Errorf("err parse msg: %w", err)
	}
	
	var hashcash HashcashData
	err = json.Unmarshal([]byte(msg.Payload), &hashcash)
	if err != nil {
		return "", fmt.Errorf("error unmarshal hashcash: %w", err)
	}
	fmt.Println("hashcash: ", hashcash)

	conf := ctx.Value(config.ConfigCtxKey).(*config.Config)
	hashcash, err = hashcash.CalCulateHashcash(conf.HashcashMaxLoop)
	if err != nil {
		return "", fmt.Errorf("error calculating hashcash: %w", err)
	}
	fmt.Println("hashcash calculated:", hashcash)

	byteData, err := json.Marshal(hashcash)
	if err != nil {
		return "", fmt.Errorf("err marshal hashcash: %w", err)
	}

	err = sendClientMessage(Message{
		Header:  RequestResource,
		Payload: string(byteData),
	}, writeConn)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	fmt.Println("request sent to server")

	msgStr, err = readConnMsg(reader)
	if err != nil {
		return "", fmt.Errorf("err read msg: %w", err)
	}

	msg, err = ParseMsg(msgStr)
	if err != nil {
		return "", fmt.Errorf("err parse msg: %w", err)
	}
	return msg.Payload, nil
}

func readConnMsg(reader *bufio.Reader) (string, error) {
	return reader.ReadString('\n')
}