package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/devfullcycle/20-CleanArch/pkg/events"
	"github.com/streadway/amqp"
)

type GetAllOrdersHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewGetAllOrdersHandler(rabbitMQChannel *amqp.Channel) *GetAllOrdersHandler {
	return &GetAllOrdersHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *GetAllOrdersHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Get All Orders: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
