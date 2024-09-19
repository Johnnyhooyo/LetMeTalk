package main

import (
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
)

func init() {
	defaultServer.AddHandler(reflect.TypeOf(&messages.SayRequest{}), func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.SayRequest)
		notifyAll(context, &messages.SayResponse{
			UserName: msg.UserName,
			Message:  msg.Message,
		})
	})
}
