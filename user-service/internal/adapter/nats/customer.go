package producer

import (
	"context"
	"fmt"
	"github.com/AskatNa/FoodStore/user-service/internal/adapter/nats/dto"
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"github.com/AskatNa/FoodStore/user-service/pkg/nats"
	"log"
	"time"

	"google.golang.org/protobuf/proto"
)

const PushTimeout = time.Second * 30

type Customer struct {
	natsClient *nats.Client
	subject    string
}

func NewCustomerProducer(
	natsClient *nats.Client,
	subject string,
) *Customer {
	return &Customer{
		natsClient: natsClient,
		subject:    subject,
	}
}

func (c *Customer) Push(ctx context.Context, customer model.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, PushTimeout)
	defer cancel()

	pbCustomer := dto.FromCustomer(customer)
	data, err := proto.Marshal(pbCustomer)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}

	err = c.natsClient.Conn.Publish(c.subject, data)
	if err != nil {
		return fmt.Errorf("c.natsClient.Conn.Publish: %w", err)
	}
	log.Println("customer is pushed:", customer)

	return nil
}
