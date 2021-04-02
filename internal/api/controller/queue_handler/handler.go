package queue_handler

import (
	"golang-gin-api/pkg/rabbitmq"
	"net/http"
)

var _ MqController = (*mqController)(nil)

var (
	QueueEntity     = rabbitmq.GetQueueEntity()
	RabbitMQ        = rabbitmq.GetRabbitMQ()
	QueueBindEntity = rabbitmq.GetQueueBindEntity()
	MessageEntity   = rabbitmq.GetMessageEntity()
	ExchangeEntity  = rabbitmq.GetExchangeEntity()
)

type MqController interface {
	QueueBindHandler(w http.ResponseWriter, r *http.Request)
	QueueHandler(w http.ResponseWriter, r *http.Request)
	PublishHandler(w http.ResponseWriter, r *http.Request)
	ExchangeHandler(w http.ResponseWriter, r *http.Request)
}

type mqController struct {
}

func MqHandler() MqController {
	return &mqController{}
}
