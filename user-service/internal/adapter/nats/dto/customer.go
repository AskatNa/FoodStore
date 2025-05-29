package dto

import (
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"github.com/AskatNa/apis-gen-user-service/events/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromCustomer(client model.Customer) *events.Customer {
	return &events.Customer{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		UpdatedAt: timestamppb.New(client.UpdatedAt),
		IsDeleted: client.IsDeleted,
	}
}
