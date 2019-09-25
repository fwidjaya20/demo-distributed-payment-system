package config

const (
	SERVICE_NAME = "SERVICE_NAME"
	NATS_CLIENT_ID = "NATS_CLIENT_ID"
)

const (
	EVENT_STORE_NATS_CLUSTER = "EVENT_STORE_NATS_CLUSTER"
	EVENT_STORE_NATS_ADDR = "EVENT_STORE_NATS_ADDR"
)

var defaultConfig = map[string]string {
	SERVICE_NAME:				"Payment System",
	EVENT_STORE_NATS_ADDR: 		"nats://localhost:4222",
	EVENT_STORE_NATS_CLUSTER:	"test-cluster",
	NATS_CLIENT_ID:				"payment-system-api",
}

func GetEnv(key string) string {
	return defaultConfig[key]
}
