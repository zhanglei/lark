package global

type KafkaMessageHandler func(msg []byte, msgKey string) (err error)
