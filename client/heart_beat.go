package main

import (
	"github.com.johnyooho.lmt/common"
	messages "github.com.johnyooho.lmt/messages"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
)

func init() {

	handlerMap[reflect.TypeOf(&messages.HeartBeatReq{})] = func(context actor.Context, i interface{}) {
		common.DefaultLogger.Debug("ping from server")
		client.RootContext.Send(client.Server, &messages.HeartBeatResp{
			Pong:   1,
			Sender: client.PID,
		})
	}
}
