package internal

import (
	"reflect"
	"server/msg"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	register(&msg.Hello{}, onHello)
}

func register(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func onHello(args []interface{}) {
	m := args[0].(*msg.Hello)
	agent := args[1].(gate.Agent)

	log.Debug("hello %v", m.Name)

	// resp with a `hello`
	respMsg := &msg.Hello{
		Name: "client",
	}
	agent.WriteMsg(respMsg)
}
