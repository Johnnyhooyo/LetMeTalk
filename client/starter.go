package main

import (
	"github.com.johnyooho.lmt/common"
	"github.com/asynkron/protoactor-go/actor"
	"reflect"
)

func init() {
	handlerMap[reflect.TypeOf(&actor.Started{})] = func(context actor.Context, i interface{}) {
		common.DefaultLogger.Info("server started")
	}
}
