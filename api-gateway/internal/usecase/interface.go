package usecase

import (
	"context"
	"github.com/AskatNa/FoodStore/api-gateway/internal/model"
)

type ClientPresenter interface {
	Create(ctx context.Context, request model.Client) (model.Client, error)
	Update(ctx context.Context, request model.Client) (model.Client, error)
	Get(ctx context.Context, id uint64) (model.Client, error)
	Delete(ctx context.Context, id uint64) error
}
