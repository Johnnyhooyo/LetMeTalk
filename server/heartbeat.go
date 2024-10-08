package main

import (
	"github.com.johnyooho.lmt/common"
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
	"time"
)

var pingClients = actor.NewPIDSet()

func init() {
	defaultServer.AddHandler(reflect.TypeOf(&messages.HeartBeatResp{}), func(context actor.Context, i interface{}) {
		msg, _ := i.(*messages.HeartBeatResp)
		common.DefaultLogger.Debug("pong from client %v", msg.Sender)
		removed := pingClients.Remove(msg.Sender)
		common.DefaultLogger.Debug("remove result is %t", removed)
	})
}

func MonitorClients(clients *actor.PIDSet, rootContext *actor.RootContext) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			for _, client := range pingClients.Values() {
				common.DefaultLogger.Info("Client %v disConnected", client)
			}
			for _, client := range clients.Values() {
				rootContext.Send(client, &messages.HeartBeatReq{Ping: 1})
				pingClients.Add(client)
			}
		}
	}()
}
