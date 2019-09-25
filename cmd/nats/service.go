package nats

import (
	"fmt"
	"github.com/fwidjaya20/demo-distributed-payment-system/config"
	"github.com/nats-io/go-nats"
	stan "github.com/nats-io/go-nats-streaming"
	"log"
	"os"
)

func InitNatsConnection() stan.Conn {
	cluster := config.GetEnv(config.EVENT_STORE_NATS_CLUSTER)
	natsClientId := config.GetEnv(config.NATS_CLIENT_ID)
	natsAddr := config.GetEnv(config.EVENT_STORE_NATS_ADDR)

	fmt.Println(cluster, natsClientId, natsAddr)

	nc, err := nats.Connect(natsAddr)

	if nil != err {
		log.Println("transport", "nats", "err", err)
		os.Exit(1)
	}

	natsConn, err := stan.Connect(cluster, natsClientId, stan.NatsConn(nc))

	if nil != err {
		log.Println("transport", "nats", "err", err)
		os.Exit(1)
	}

	log.Println("transport", "nats", "addr", natsAddr)

	return natsConn
}
