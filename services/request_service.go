package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/megamsquare/word-of-wisdom/models"
	"github.com/megamsquare/word-of-wisdom/pkg/config"
)

const (
	Quit = iota
	RequestChallenge
	ResponseChallenge
	RequestResource
	ResponseResource
)

func ProcessRequest(ctx context.Context, messageString string, clientInfo string) (*Message, error) {
	msg, err := ParseMsg(messageString)
	if err != nil {
		return nil, err
	}

	switch msg.Header {
	case Quit:
		return nil, errors.New("client close connection")
	case RequestChallenge:
		fmt.Printf("%s requests for challenge\n", clientInfo)

		conf, ok := ctx.Value(config.ConfigCtxKey).(*config.Config)
		if !ok {
			fmt.Println("Config not ok: ", conf)
		}

		hashcash := HashcashData{
			Version:    1,
			ZerosCount: conf.HashcashZerosCount,
			Date:       time.Now().Unix(),
			Resource:   clientInfo,
			Rand:       base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", rand.Intn(100000)))),
			Counter:    0,
		}
		hashcashMarshal, err := json.Marshal(hashcash)
		if err != nil {
			return nil, fmt.Errorf("hashcash error marshal: %v", err)
		}
		msg := Message{
			Header:  ResponseChallenge,
			Payload: string(hashcashMarshal),
		}
		return &msg, nil
	case RequestResource:
		fmt.Printf("%s requests for resource\n", clientInfo)

		var hashcash HashcashData
		err = json.Unmarshal([]byte(msg.Payload), &hashcash)
		if err != nil {
			return nil, fmt.Errorf("requested data is not valid JSON: %v", err)
		}

		if hashcash.Resource != clientInfo {
			return nil, fmt.Errorf("resource name mismatch: got '%s', want '%s'", hashcash.Resource, clientInfo)
		}

		conf := ctx.Value(config.ConfigCtxKey).(*config.Config)

		if time.Now().Unix()-hashcash.Date > conf.HashcashDuration {
			return nil, fmt.Errorf("date expired")
		}

		maxLoop := hashcash.Counter
		if maxLoop == 0 {
			maxLoop = 1
		}

		_, err = hashcash.CalCulateHashcash(maxLoop)
		if err != nil {
			return nil, fmt.Errorf("calculate hashcash failed: %v", err)
		}

		fmt.Printf("%s computed hashcash succesfully %s\n", clientInfo, msg.Payload)
		msg := Message{
			Header:  ResponseResource,
			Payload: models.Quotes[rand.Intn(4)],
		}

		return &msg, nil
	default:
		return nil, fmt.Errorf("invalid header")
	}
}
