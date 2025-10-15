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
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	user.Uuid = uuid.New()
	query := `INSERT INTO users (uuid, name, email, password, contact, address, status) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING uuid, name, email, password, contact, address, status`
	row := tx.QueryRowContext(ctx, query, user.Uuid, user.Name, user.Email, user.Password, user.Pnone, user.Address, user.Status)

	var newUser models.Users
	err = row.Scan(&newUser.Uuid, &newUser.Name, &newUser.Email, &newUser.Password, &newUser.Pnone, &newUser.Address, &newUser.Status)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			r.logger.Errorf("tx rollback error: %v", rbErr)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			r.logger.Errorf("tx rollback error: %v", rbErr)
		}
		return nil, err
	}

	r.logger.Infof("User created with UUID: %s", newUser.Uuid)
	return &newUser, nil
}
