package server

import "github.com/AskatNa/FoodStore/user-service/internal/adapter/gRPC/frontend"

type CustomerUseCase interface {
	frontend.CustomerUseCase
}
