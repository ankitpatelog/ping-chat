package handler

import (
	"ping/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

// constructor
func NewUserservice(pool *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		handler: service.NewUserservice(pool),
	}
}

type UserHandler struct{
	handler *service.UserService
}

func ()CreateUser()  {
	
}
