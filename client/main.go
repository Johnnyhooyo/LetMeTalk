package main

import (
	messages "github.com.johnyooho.lmt/messages"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"log"
	"reflect"
)

type Client struct {
	RootContext *actor.RootContext
	Server      *actor.PID
	PID         *actor.PID
	handlers    map[reflect.Type]Handler
}

var (
	client *Client
	Nick   = "Roger"
	cons   = console.NewConsole(func(text string) {
		client.Send(&messages.SayRequest{UserName: Nick, Message: text})
	})
	handlerMap = make(map[reflect.Type]Handler)
)

type Handler func(actor.Context, interface{})

func (c2 *Client) AddHandler(msgType reflect.Type, handler Handler) {
	// 添加新的消息类型及其处理函数
	c2.handlers[msgType] = handler
}

func (c2 *Client) Receive(ctx actor.Context) {
	msg := ctx.Message()
	msgType := reflect.TypeOf(msg) // 获取消息的类型

	// 查找是否存在对应消息类型的处理函数
	if handler, ok := c2.handlers[msgType]; ok {
		handler(ctx, msg) // 调用处理函数
	} else {
		log.Printf("no handler found for message type: %T", msg)
	}
}

func (c2 *Client) Send(msg any) {
	c2.RootContext.Send(c2.Server, msg)
}

func main() {
	system := actor.NewActorSystem()
	config := remote.Configure("127.0.0.1", 0)
	remoter := remote.NewRemote(system, config)
	remoter.Start()

	server := actor.NewPID("127.0.0.1:8080", "chatserver")

	// define root context
	rootContext := system.Root

	client = &Client{
		RootContext: rootContext,
		Server:      server,
		handlers:    handlerMap,
	}
	// spawn our chat client inline
	props := actor.PropsFromProducer(func() actor.Actor { return client })

	clientPid := rootContext.Spawn(props)
	client.PID = clientPid

	rootContext.Send(server, &messages.Connect{
		Sender: clientPid,
		Name:   Nick,
	})

	cons.Run()

}
