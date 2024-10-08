package main

import (
	"fmt"
	"github.com.johnyooho.lmt/common"
	messages "github.com.johnyooho.lmt/messages"
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"log/slog"
	"os"
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
		common.DefaultLogger.Error("no handler found for message type: %T", msg)
	}
}

func (c2 *Client) Send(msg any) {
	c2.RootContext.Send(c2.Server, msg)
}

func main() {
	args := os.Args
	addr := ":8091"
	if len(args) == 4 {
		ip := args[1]
		port := args[2]
		Nick = args[3]
		addr = fmt.Sprintf("%s:%s", ip, port)
	}
	system := actor.NewActorSystem(actor.WithLoggerFactory(func(system *actor.ActorSystem) *slog.Logger {
		return common.DebugLogger
	}))
	config := remote.Configure("", 8092, remote.WithAdvertisedHost("localhost:8092"))
	remoter := remote.NewRemote(system, config)
	remoter.Start()

	server := actor.NewPID(addr, "chatserver")

	// define root context
	rootContext := system.Root.WithSenderMiddleware(func(next actor.SenderFunc) actor.SenderFunc {
		return func(ctx actor.SenderContext, target *actor.PID, envelope *actor.MessageEnvelope) {
			common.DebugLogger.Debug("Actor send message", slog.Any("self", ctx.Self()), slog.Any("target", ctx.Sender()), slog.Any("message", ctx.Message()))
			next(ctx, target, envelope)
		}
	})

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
