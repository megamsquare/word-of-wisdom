package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/megamsquare/word-of-wisdom/pkg/config"
	"github.com/megamsquare/word-of-wisdom/services"
)

func main() {
	fmt.Println("start server")

	configInst, err := config.Load("env/env.json")
	if err != nil {
		fmt.Println("loading config error: ", err)
		return
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, config.ConfigCtxKey, configInst)

	rand.Seed(time.Now().UnixNano())

	serverAddress := fmt.Sprintf("%s:%d", configInst.ServerHost, configInst.ServerPort)
	err = services.RunServer(ctx, serverAddress)
	if err != nil {
		fmt.Println("error with server", err)
	}

}
