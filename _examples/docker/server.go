package main

import (
	"time"

	"github.com/saiser/rpcx/log"
	"github.com/saiser/rpcx"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

type Arith2 int

func (t *Arith2) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B * 10
	return nil
}

func main() {
	server1 := rpcx.NewServer()
	server1.RegisterName("Arith", new(Arith))
	server1.Start("tcp", ":8972")

	time.Sleep(5 * time.Second)
	log.Info(server1.Address())

	server2 := rpcx.NewServer()
	server2.RegisterName("Arith", new(Arith2))
	server2.Serve("tcp", ":8973")
}
