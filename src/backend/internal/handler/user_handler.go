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

type UserHandler struct {
	handler *service.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.CreateUserRequest
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// convert DTO → model
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userReturned, err := h.handler.CreateProcess(ctx, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userReturned)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req models.LoginCred
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	// call login service
	lgRes, err := h.handler.LoginProcess(ctx, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// directly return service response
	json.NewEncoder(w).Encode(lgRes)

}
