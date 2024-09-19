package main

import (
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"reflect"
)

func init() {

	handlerMap[reflect.TypeOf(&messages.SayResponse{})] = func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.SayResponse)
		log.Printf("%v: %v", msg.UserName, msg.Message)
	}
}
