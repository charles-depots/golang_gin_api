package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"golang-gin-api/config"
)

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
	MqUrl     string
}

// Init rabbitmq client
func InitRabbitMQ(queueName string, exchange string, key string) (*RabbitMQ, error) {
	rabbitMQ := config.GetConfig().RabbitMQ
	MQURL := fmt.Sprintf("amqp://%s:%s@%s/%s", rabbitMQ.User, rabbitMQ.Passwd, rabbitMQ.Addr, rabbitMQ.Vhost)
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: MQURL}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "An error occurred while establishing the connection")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "Unable to get channel")
	return rabbitmq, err
}

// Error handle
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}
