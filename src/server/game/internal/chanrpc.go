package internal

import (
	"fmt"

	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	fmt.Printf("rpcNewAgent, a = %+v\n", rpcNewAgent, a)
	_ = a
}

func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	fmt.Printf("rpcCloseAgent, a = %+v\n", rpcCloseAgent, a)
	_ = a
}
