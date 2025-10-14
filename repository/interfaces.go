package repository

import (
	"context"
	"pov_golang/models"
)

type UserRepository interface {
	Create(ctx context.Context, req models.Users) (*models.Users, error)
}
