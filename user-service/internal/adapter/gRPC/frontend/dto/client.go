package dto

import (
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	base "github.com/AskatNa/apis-gen-user-service/base/frontend/v1"
	svc "github.com/AskatNa/apis-gen-user-service/service/frontend/client/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCustomerFromRegisterRequest(req *svc.RegisterRequest) (model.Customer, error) {
	return model.Customer{
		Email:       req.Email,
		NewPassword: req.Password,
	}, nil
}

func ToCustomerFromUpdateRequest(req *svc.UpdateRequest) (model.Customer, error) {
	return model.Customer{
		ID:              req.Id,
		Name:            req.Name,
		Phone:           req.Phone,
		Email:           req.Email,
		CurrentPassword: req.OldPassword,
		NewPassword:     req.Password,
	}, nil
}

func FromCustomer(client model.Customer) *base.Client {
	return &base.Client{
		Id:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		Phone:     client.Phone,
		CreatedAt: timestamppb.New(client.CreatedAt),
		//UpdatedAt: timestamppb.New(client.UpdatedAt),
		IsDeleted: client.IsDeleted,
	}
}
