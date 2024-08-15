package services

import (
	"errors"
	"os"
	"time"

	"github.com/amrimuf/hompimRent/models"
	"github.com/amrimuf/hompimRent/repositories"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    Repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
    return &AuthService{Repo: repo}
}

func (s *AuthService) Register(user *models.User, password string) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    user.PasswordHash = string(hashedPassword)
    user.ID = uuid.Must(uuid.NewV4()) 
    user.Role = "user"                

    return s.Repo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {
    user, err := s.Repo.FindByEmail(email)
    if err != nil {
        return "", errors.New("user not found")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
        return "", errors.New("invalid credentials")
    }

    token, err := s.generateJWT(user)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (s *AuthService) generateJWT(user *models.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID.String(),
        "email":   user.Email,
        "role":    user.Role,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })

    secretKey := os.Getenv("JWT_SECRET")

    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}
