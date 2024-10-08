package main

import (
	"fmt"
	"github.com.johnyooho.lmt/common"
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
)

func init() {
	defaultServer.AddHandler(reflect.TypeOf(&messages.Connect{}), func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.Connect)
		common.DefaultLogger.Debug("Client %v connected", msg.Sender)
		defaultServer.AddClient(msg.Sender)
		context.Send(msg.Sender, &messages.Connected{Message: fmt.Sprintf("Welcome %s!", msg.Name)})
	})
}
