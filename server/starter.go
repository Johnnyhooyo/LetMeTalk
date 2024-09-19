package main

import (
	"github.com/asynkron/protoactor-go/actor"
	"log"
	"reflect"
)

func init() {
	defaultServer.AddHandler(reflect.TypeOf(&actor.Started{}), func(context actor.Context, i interface{}) {
		log.Printf("server started")
	})
}
