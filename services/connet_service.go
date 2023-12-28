package services

import (
	"context"
	"fmt"
	"net"
	"time"
)

func RunServer(ctx context.Context, addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	defer listener.Close()
	fmt.Println("server listening ", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			return fmt.Errorf("error with connection: %v", err)
		}

		go handleServerConnection(ctx, conn)
	}
}

func RunClient(ctx context.Context, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}

	defer conn.Close()
	fmt.Println("client connect to", addr)

	for {
		msg, err :=handleClientConnection(ctx, conn, conn)
		if err != nil {
			return err
		}
		fmt.Println("quote result: ", msg)
		time.Sleep(10 * time.Second)
	}
}