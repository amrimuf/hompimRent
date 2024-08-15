package services

import (
	"github.com/amrimuf/hompimRent/models"
	"github.com/amrimuf/hompimRent/repositories"
	"github.com/gofrs/uuid"
)

type UserService struct {
    Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{Repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
    return s.Repo.Create(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetAll()
}

func (s *UserService) GetUserByID(id uuid.UUID) (*models.User, error) {
    return s.Repo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
    return s.Repo.Update(user)
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
    return s.Repo.Delete(id)
}
