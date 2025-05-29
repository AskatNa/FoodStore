package dto

import (
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"github.com/AskatNa/apis-gen-user-service/events/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromCustomer(client model.Customer) *events.Client {
	return &events.Client{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		//UpdatedAt: timestamppb.New(client.UpdatedAt),
		IsDeleted: client.IsDeleted,
	}
}
