package main

import (
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"reflect"
)

func init() {

	handlerMap[reflect.TypeOf(&messages.Connected{})] = func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.Connected)
		log.Println(msg.Message)
	}
}
