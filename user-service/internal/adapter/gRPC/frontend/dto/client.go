package dto

import (
	"fmt"
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	base "github.com/AskatNa/apis-gen-user-service/base/frontend/v1"
	svc "github.com/AskatNa/apis-gen-user-service/service/frontend/client/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCustomerFromRegisterRequest(req *svc.RegisterRequest) (model.Customer, error) {
	if req.Email == "" || req.Password == "" {
		return model.Customer{}, fmt.Errorf("email and password are required for registration")
	}

	return model.Customer{
		Email:       req.Email,
		NewPassword: req.Password,
	}, nil
}

func ToCustomerFromUpdateRequest(req *svc.UpdateRequest) (model.Customer, error) {
	customer := model.Customer{
		ID: req.Id,
	}

	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.Email != "" {
		customer.Email = req.Email
	}
	if req.Password != "" {
		customer.NewPassword = req.Password
	}
	if req.OldPassword != "" {
		customer.CurrentPassword = req.OldPassword
	}

	return customer, nil
}

func FromCustomer(client model.Customer) *base.Customer {
	return &base.Customer{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		UpdatedAt: timestamppb.New(client.UpdatedAt),
		IsDeleted: client.IsDeleted,
	}
}
