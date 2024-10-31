package queue

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"study/application/model"
	"study/infrastructure/parser"
	"time"
)

type Rabbit struct {
	conn     *amqp.Connection
	messages <-chan amqp.Delivery
}

func NewRabbit() *Rabbit {
	conn, err := amqp.DialConfig("amqp://admin:123456@localhost:5672", amqp.Config{
		Heartbeat: 60 * time.Second,
	})
	if err != nil {
		fmt.Println("~[Rabbit] error connection", err.Error())
		return nil
	}
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println("~[Rabbit] error channel", err.Error())
		return nil
	}
	m, err := ch.Consume(
		"sync_camera",
		"", false, false, false, false, nil,
	)
	if err != nil {
		fmt.Println("~[Rabbit] error consume", err.Error())
		return nil
	}
	return &Rabbit{
		conn:     conn,
		messages: m,
	}
}

func (r *Rabbit) Listen(callback func(d []*model.Camera)) {
	for d := range r.messages {
		fmt.Printf("~[Rabbit] received a message: %s\n", d.Body)
		p := parser.NewParser()
		var data []*model.Camera
		if err := p.Decode(d.Body, &data); err != nil {
			fmt.Println("~[Rabbit] error im parser message", err)
			return
		}
		callback(data)
		if err := d.Ack(false); err != nil {
			fmt.Printf("~[Rabbit] error in ack message!")
			continue
		}
	}
}
