package dao

import (
	"github.com/AskatNa/FoodStore/user-service/internal/model"
	"time"
)

type Session struct {
	UserID       uint64    `bson:"userID"`
	RefreshToken string    `bson:"refreshToken"`
	ExpiresAt    time.Time `bson:"expiresAt"`
	CreatedAt    time.Time `bson:"createdAt"`
}

func FromSession(session model.Session) Session {
	return Session{
		UserID:       session.UserID,
		RefreshToken: session.RefreshToken,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}
}

func ToSession(daoSession Session) model.Session {
	return model.Session{
		UserID:       daoSession.UserID,
		RefreshToken: daoSession.RefreshToken,
		ExpiresAt:    daoSession.ExpiresAt,
		CreatedAt:    daoSession.CreatedAt,
	}
}
