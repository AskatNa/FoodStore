package service

import (
	"github.com/AskatNa/FoodStore/api-gateway/internal/adapter/http/server/handler"
)

type ClientUsecase interface {
	handler.ClientUsecase
}
