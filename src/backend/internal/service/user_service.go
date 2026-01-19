package service

import (
	"context"
	"fmt"
	"math/rand/v2"
	"ping/internal/models"
	"ping/internal/repo"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
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

func (s *UserService) CreateProcess(ctx context.Context, currUser *models.User) (models.User, error) {
    // first check user in db
    exists, err := s.service.CheckUser(ctx, currUser.Email)
    if err != nil {
        return models.User{}, err // Return empty user and error
    }

    if exists {
        // If user exists, return an error or the existing user
        // Returning an error here is standard for a "Create" process
        return models.User{}, fmt.Errorf("user with email %s already exists", currUser.Email)
    }

    // for new user
    // give unique id
    currUser.ID = uuid.NewString()

    // generate hashed pass
    pass, err := bcrypt.GenerateFromPassword([]byte(currUser.Password), 10)
    if err != nil {
        return models.User{}, err
    }

    currUser.Password = string(pass)

    // generate unique username
    newusername, err := s.GenUniqUsername(ctx, currUser.Name)
    if err != nil {
        return models.User{}, err
    }
    currUser.Username = newusername

    // now call save handler 
    user, err := s.service.SaveUser(ctx, *currUser)
    if err != nil {
        return models.User{}, err
    }

    // Return the saved user object and nil for error
    return user, nil
}

//func for unique username generation
func (s *UserService) GenUniqUsername(ctx context.Context, name string) (string, error) {
    lowerName := strings.ToLower(strings.ReplaceAll(name, " ", "")) // Remove spaces & lowercase

    for { // Infinite loop until we find a unique name or hit an error
        // Generate random 4-digit number (1000-9999)
        randomNum := rand.IntN(9000) + 1000
        username := fmt.Sprintf("%s%d", lowerName, randomNum)

        // Check if it exists in the DB
        isExist, err := s.service.CheckUsername(ctx, username)
        if err != nil {
            return "", err // Database/Network error
        }

        if !isExist {
            // Success: Username is unique
            return username, nil
        }
        
        // If isExist is true, the loop continues to try a new number
    }
}


