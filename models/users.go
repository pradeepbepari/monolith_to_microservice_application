package models

import "github.com/google/uuid"

type Users struct {
	Uuid     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Pnone    int64     `json:"contact"`
	Address  string    `json:"address"`
	Status   bool      `json:"status"`
}
