package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"golang-gin-api/config"
	"log"
)

// Entity for HTTP Request Body: Message/Exchange/Queue/QueueBind JSON Input
type MessageEntity struct {
	Exchange     string `json:"exchange"`
	Key          string `json:"key"`
	DeliveryMode uint8  `json:"deliverymode"`
	Priority     uint8  `json:"priority"`
	Body         string `json:"body"`
}

type ExchangeEntity struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	NoWait     bool   `json:"nowait"`
}

type QueueEntity struct {
	Name       string `json:"name"`
	Durable    bool   `json:"durable"`
	AutoDelete bool   `json:"autodelete"`
	Exclusive  bool   `json:"exclusive"`
	NoWait     bool   `json:"nowait"`
}

type QueueBindEntity struct {
	Queue    string   `json:"queue"`
	Exchange string   `json:"exchange"`
	NoWait   bool     `json:"nowait"`
	Keys     []string `json:"keys"` // bind/routing keys
}

// RabbitMQ Operate Wrapper
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan error
}

func GetQueueEntity() *QueueEntity {
	return new(QueueEntity)
}

func GetQueueBindEntity() *QueueBindEntity {
	return new(QueueBindEntity)
}

func GetRabbitMQ() *RabbitMQ {
	return new(RabbitMQ)
}

func GetMessageEntity() *MessageEntity {
	return new(MessageEntity)
}

func GetExchangeEntity() *ExchangeEntity {
	return new(ExchangeEntity)
}

func (r *RabbitMQ) Connect() (err error) {
	rabbitMQ := config.GetConfig().RabbitMQ
	amqpUri := fmt.Sprintf("amqp://%s:%s@%s/", rabbitMQ.User, rabbitMQ.Passwd, rabbitMQ.Addr)
	r.conn, err = amqp.Dial(amqpUri)
	if err != nil {
		log.Printf("[amqp] connect error: %s\n", err)
		return err
	}
	r.channel, err = r.conn.Channel()
	if err != nil {
		log.Printf("[amqp] get channel error: %s\n", err)
		return err
	}
	r.done = make(chan error)
	return nil
}

// Publish sends a Publishing from the client to an exchange on the server.
func (r *RabbitMQ) Publish(exchange, key string, deliverymode, priority uint8, body string) (err error) {
	err = r.channel.Publish(exchange, key, false, false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    deliverymode,
			Priority:        priority,
			Body:            []byte(body),
		},
	)
	if err != nil {
		log.Printf("[amqp] publish message error: %s\n", err)
		return err
	}
	return nil
}

// ExchangeDeclare declares an exchange on the server.
func (r *RabbitMQ) DeclareExchange(name, typ string, durable, autodelete, nowait bool) (err error) {
	err = r.channel.ExchangeDeclare(name, typ, durable, autodelete, false, nowait, nil)
	if err != nil {
		log.Printf("[amqp] declare exchange error: %s\n", err)
		return err
	}
	return nil
}

// ExchangeDelete removes the named exchange from the server.
func (r *RabbitMQ) DeleteExchange(name string) (err error) {
	err = r.channel.ExchangeDelete(name, false, false)
	if err != nil {
		log.Printf("[amqp] delete exchange error: %s\n", err)
		return err
	}
	return nil
}

// QueueDeclare declares a queue to hold messages and deliver to consumers.
func (r *RabbitMQ) DeclareQueue(name string, durable, autodelete, exclusive, nowait bool) (err error) {
	_, err = r.channel.QueueDeclare(name, durable, autodelete, exclusive, nowait, nil)
	if err != nil {
		log.Printf("[amqp] declare queue error: %s\n", err)
		return err
	}
	return nil
}

// QueueDelete removes the queue from the server.
func (r *RabbitMQ) DeleteQueue(name string) (err error) {
	// TODO: other property wrapper
	_, err = r.channel.QueueDelete(name, false, false, false)
	if err != nil {
		log.Printf("[amqp] delete queue error: %s\n", err)
		return err
	}
	return nil
}

// QueueBind binds an exchange to a queue
func (r *RabbitMQ) BindQueue(queue, exchange string, keys []string, nowait bool) (err error) {
	for _, key := range keys {
		if err = r.channel.QueueBind(queue, key, exchange, nowait, nil); err != nil {
			log.Printf("[amqp] bind queue error: %s\n", err)
			return err
		}
	}
	return nil
}

// QueueUnbind removes a binding between an exchange and queue matching the key and arguments.
func (r *RabbitMQ) UnBindQueue(queue, exchange string, keys []string) (err error) {
	for _, key := range keys {
		if err = r.channel.QueueUnbind(queue, key, exchange, nil); err != nil {
			log.Printf("[amqp] unbind queue error: %s\n", err)
			return err
		}
	}
	return nil
}

// Consume immediately starts delivering queued messages.
func (r *RabbitMQ) ConsumeQueue(queue string, message chan []byte) (err error) {
	deliveries, err := r.channel.Consume(queue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("[amqp] consume queue error: %s\n", err)
		return err
	}
	go func(deliveries <-chan amqp.Delivery, done chan error, message chan []byte) {
		for d := range deliveries {
			message <- d.Body
		}
		done <- nil
	}(deliveries, r.done, message)
	return nil
}

// Close requests and waits for the response to close the AMQP connection.
func (r *RabbitMQ) Close() (err error) {
	err = r.conn.Close()
	if err != nil {
		log.Printf("[amqp] close error: %s\n", err)
		return err
	}
	return nil
}
