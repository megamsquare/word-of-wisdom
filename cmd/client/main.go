package main

import (
	"context"
	"fmt"

	"github.com/megamsquare/word-of-wisdom/pkg/config"
	"github.com/megamsquare/word-of-wisdom/services"
)

func main()  {
	fmt.Println("start client")

	configInst, err := config.Load("env/env.json")
	if err != nil {
		fmt.Println("loading config error: ", err)
		return
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, config.ConfigCtxKey, configInst)

	serverAddress := fmt.Sprintf("%s:%d", configInst.ServerHost, configInst.ServerPort)
	err = services.RunClient(ctx, serverAddress)
	if err != nil {
		fmt.Println("error with client", err)
	}

}