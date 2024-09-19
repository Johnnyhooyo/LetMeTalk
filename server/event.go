package main

import (
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
)

func init() {
	defaultServer.AddHandler(reflect.TypeOf(&messages.NickRequest{}), func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.NickRequest)
		notifyAll(context, &messages.NickResponse{
			OldUserName: msg.OldUserName,
			NewUserName: msg.NewUserName,
		})
	})
}
