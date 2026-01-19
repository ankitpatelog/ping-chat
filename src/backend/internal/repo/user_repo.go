package repo

import (
	"context"

	"ping/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

// constructor
func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		pool: pool,
	}
}

// CreateUser inserts a new user into ping.users
func (r *UserRepo) SaveUser(ctx context.Context, user models.User) (models.User, error) {

	query := `
		INSERT INTO ping.users (id, name, username, email, password)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		user.ID,        
		user.Name,
		user.Username,
		user.Email,
		user.Password,
	)

	if err != nil {
		return models.User{}, err
	}

	// created_at / updated_at are set by DB defaults
	return user, nil
}
