package main

import (
	"fmt"
	"github.com/fwidjaya20/demo-distributed-payment-system/cmd/nats"
	stan "github.com/nats-io/go-nats-streaming"
	"log"
	"runtime"
	"time"
)

func main() {
	natsConn := nats.InitNatsConnection()

	aw, _ := time.ParseDuration("60s")
	s, _ := natsConn.Subscribe("OrderSystem.Order", func(msg *stan.Msg) {
		_ = msg.Ack()

		log.Println("Message", msg.Data)

		go func() {
			fmt.Println("Sleep Time")
			time.Sleep(10 * time.Second)
		}()
	},
		stan.DurableName("payment-system-durable"),
		stan.SetManualAckMode(),
		stan.AckWait(aw),
	)

	_ = s.Unsubscribe()
	runtime.Goexit()
}
