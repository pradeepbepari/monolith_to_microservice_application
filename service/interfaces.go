package service

import (
	"context"
	"pov_golang/models"
)

type UserService interface {
	Create(ctx context.Context, req models.Users) (*models.Users, error)
}
