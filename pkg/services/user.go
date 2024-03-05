package services

import (
	"HarlanCinema/pkg/models"
	repo "HarlanCinema/pkg/repos"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repo.Repository
}

func NewUserService(Repo *repo.Repository) *UserService {
	return &UserService{Repo: Repo}
}

func (s *UserService) RegisterUser(user models.User) (models.User, error) {
	user.ID = uuid.New().String()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hash)
	user.Role = "user"
	user, err = s.Repo.UserRepository.Create(user)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) AuthenticateUser(login, pass string) (models.User, error) {

	user, err := s.Repo.UserRepository.GetUserByUsername(login)
	if err != nil {
		return models.User{}, models.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return models.User{}, models.ErrInvalidCredentials
	}

	return user, nil
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	user, err := s.Repo.UserRepository.GetUserByID(id)
	if err != nil {
		return models.User{}, models.NotFoundAnything
	}
	return user, nil
}
