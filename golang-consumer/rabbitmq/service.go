package rabbitmq

import (
	"golang-consumer/config"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Conn struct {
	Channel *amqp.Channel
}

func New() (*Conn, error) {
	AMQP_URI := config.New().AmqpUri

	con, err := amqp.DialConfig(AMQP_URI, amqp.Config{
		Heartbeat: 10 * time.Second,
	})
	if err != nil {
		defer con.Close()
		return nil, err
	}

	ch, err := con.Channel()
	if err != nil {
		defer ch.Close()
		return nil, err
	}

	ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	log.Printf("\nConnect #%s success\n", AMQP_URI)

	return &Conn{
		Channel: ch,
	}, nil
}

func (c Conn) StartCosumer(qName string, consumerName string, handleFunc func(amqp.Delivery)) error {
	q, err := c.Channel.QueueDeclare(
		qName, //name
		true,  //durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	if err != nil {
		return err
	}

	msq, err := c.Channel.Consume(
		q.Name,       // queue
		consumerName, // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)

	if err != nil {
		return err
	}

	go consumeLoop(msq, handleFunc)

	return nil
}

func consumeLoop(delivery <-chan amqp.Delivery, handleFunc func(s amqp.Delivery)) {
	for d := range delivery {
		handleFunc(d)
	}

}
