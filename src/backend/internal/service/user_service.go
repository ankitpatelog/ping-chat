package service

import (
	"ping/internal/repo"

	"github.com/jackc/pgx/v5/pgxpool"
)

// constructor
func NewUserservice(pool *pgxpool.Pool) *UserService {
	return &UserService{
		service: repo.NewUserRepo(pool),
	}
}

type UserService struct{
	service *repo.UserRepo
}