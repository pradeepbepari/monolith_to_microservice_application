package service

import (
	"context"
	"fmt"
	"log"
	"pov_golang/models"
	"pov_golang/repository"
)

type service struct {
	repository repository.UserRepository
}

func NewService(repository repository.UserRepository) repository.UserRepository {
	return service{
		repository: repository,
	}
}
func (s service) Create(ctx context.Context, req models.Users) (*models.Users, error) {
	resp, err := s.repository.Create(ctx, req)
	if err != nil {
		log.Println("ērror creating user : ", err)
		return nil, fmt.Errorf("failed to create user")
	}
	return resp, nil
}
