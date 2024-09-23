package main

import (
	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
	"log"
	"reflect"
)

func notifyAll(context actor.Context, message interface{}) {
	for _, client := range defaultServer.Clients.Values() {
		context.Send(client, message)
	}
}

var defaultServer = &Server{
	handlers: make(map[reflect.Type]Handler),
}

type Server struct {
	Clients  *actor.PIDSet
	Context  *actor.RootContext
	handlers map[reflect.Type]Handler
}

type Handler func(context actor.Context, i interface{})

func (s *Server) AddHandler(p reflect.Type, h Handler) {
	s.handlers[p] = h
}

func (s *Server) AddClient(c *actor.PID) {
	s.Clients.Add(c)
}

func (s *Server) Receive(ctx actor.Context) {
	msg := ctx.Message()
	msgType := reflect.TypeOf(msg) // 获取消息的类型

	// 查找是否存在对应消息类型的处理函数
	if handler, ok := s.handlers[msgType]; ok {
		handler(ctx, msg) // 调用处理函数
	} else {
		log.Printf("no handler found for message type: %T", msg)
	}
}

func main() {
	system := actor.NewActorSystem()
	config := remote.Configure("", 8091, remote.WithAdvertisedHost(":8091"))
	remoter := remote.NewRemote(system, config)
	remoter.Start()

	clients := actor.NewPIDSet()
	rootContext := system.Root

	defaultServer.Clients = clients
	defaultServer.Context = rootContext

	MonitorClients(clients, rootContext)

	props := actor.PropsFromProducer(func() actor.Actor { return defaultServer })

	_, _ = system.Root.SpawnNamed(props, "chatserver")
	_, _ = console.ReadLine()
}
