package main

import (
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"reflect"
)

func init() {

	handlerMap[reflect.TypeOf(&messages.HeartBeatReq{})] = func(context actor.Context, i interface{}) {
		log.Printf("ping from server")
		client.RootContext.Send(client.Server, &messages.HeartBeatResp{
			Pong:   1,
			Sender: client.PID,
		})
	}
}
