package main

import (
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"reflect"
)

func init() {

	handlerMap[reflect.TypeOf(&messages.NickResponse{})] = func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.NickResponse)
		log.Printf("%v is now known as %v", msg.OldUserName, msg.NewUserName)
	}

	cons.Command("/nick", func(newNick string) {
		Nick = newNick
		client.RootContext.Send(client.Server, &messages.NickRequest{
			OldUserName: Nick,
			NewUserName: newNick,
		})
	})
}
