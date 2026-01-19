package handler

import (
	"encoding/json"
	"net/http"
	"ping/internal/models"
	"ping/internal/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

// constructor
func NewUserHandler(pool *pgxpool.Pool) *UserHandler {
	return &UserHandler{
		handler: service.NewUserservice(pool),
	}
}

type UserHandler struct{
	handler *service.UserService
}

func (h *UserHandler)CreateUser(w http.ResponseWriter,r *http.Request)  {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	//call service handler for username and hash pass generation
	
}
