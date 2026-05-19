package user_sdk

import (
	"context"
)

type Client interface {
	GetByID(ctx context.Context, id int) (*UserWithCity, error)
	CreateUser(ctx context.Context, user *User) error
	FindWithinRadius(ctx context.Context, lat, lon, radius float64) ([]UserWithCity, error)
}
