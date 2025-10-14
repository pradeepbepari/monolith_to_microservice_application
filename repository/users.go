package repository

import (
	"context"
	"database/sql"
	"pov_golang/logger"
	"pov_golang/models"

	"github.com/google/uuid"
)

type repo struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewRepository(db *sql.DB, logger *logger.Logger) *repo {
	return &repo{
		db:     db,
		logger: logger,
	}
}
func (r *repo) Create(ctx context.Context, user models.Users) (*models.Users, error) {
	user.Uuid = uuid.New()
	query := `INSERT INTO users (uuid, name, email, password, contact, address, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING uuid, name, email, password, contact, address, status`
	row := r.db.QueryRowContext(ctx, query, user.Uuid, user.Name, user.Email, user.Password, user.Pnone, user.Address, user.Status)
	var newUser models.Users
	err := row.Scan(&newUser.Uuid, &newUser.Name, &newUser.Email, &newUser.Password, &newUser.Pnone, &newUser.Address, &newUser.Status)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
