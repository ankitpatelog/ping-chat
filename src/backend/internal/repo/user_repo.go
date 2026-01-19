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

// CreateUser inserts a new user into ping.users
func (r *UserRepo) CheckUser(ctx context.Context,email string) (bool,error) {

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM ping.users
			WHERE email=$1
		)
	`

	var exists bool

	err := r.pool.QueryRow(ctx,query,&email).Scan(&exists)
	if err!=nil {
		return false,err
	}

	return exists,nil
}

func (r *UserRepo) CheckUsername(ctx context.Context,username string) (bool,error) {
	query := `SELECT EXIST (
			SELECT 1
			FROM users
			WHERE username=$1
	)`

	var exist bool

	err := r.pool.QueryRow(ctx,query,&username).Scan(&exist)
	if err!=nil {
		return false,err
	}

	return exist,nil
}
