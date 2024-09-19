package main

import (
	"fmt"
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"reflect"
)

func init() {
	defaultServer.AddHandler(reflect.TypeOf(&messages.Connect{}), func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.Connect)
		log.Printf("Client %v connected", msg.Sender)
		defaultServer.AddClient(msg.Sender)
		context.Send(msg.Sender, &messages.Connected{Message: fmt.Sprintf("Welcome %s!", msg.Name)})
	})
}
