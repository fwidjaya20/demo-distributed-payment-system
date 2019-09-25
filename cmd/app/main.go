package main

import (
	"fmt"
	"github.com/fwidjaya20/demo-distributed-payment-system/cmd/nats"
	stan "github.com/nats-io/go-nats-streaming"
	"runtime"
)

func main() {
	natsConn := nats.InitNatsConnection()

	s, err := natsConn.Subscribe("OrderSystem.Order", func(msg *stan.Msg) {
		fmt.Println(msg.Data)
	})

	if nil != err{
		_ = s.Unsubscribe()
		panic(err)
	}

	runtime.Goexit()
}
