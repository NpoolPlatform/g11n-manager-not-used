package server

import (
	msgsrv "github.com/NpoolPlatform/go-service-framework/pkg/rabbitmq/server"
	msg "github.com/NpoolPlatform/g11n-manager/pkg/message/message"
)

func Init() error {
	return msg.InitQueues()
}

func PublishExample(example *msg.Example) error {
	return msgsrv.PublishToQueue(msg.QueueExample, example)
}
