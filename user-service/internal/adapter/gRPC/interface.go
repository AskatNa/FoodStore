package server

import "github.com/AskatNa/FoodStore/user-service/internal/adapter/gRPC/frontend"

type CustomerUsecase interface {
	frontend.CustomerUsecase
}
