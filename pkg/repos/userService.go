package repositories

import (
	"HarlanCinema/pkg/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (s *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if result := s.db.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (s *UserRepository) RegisterUser(user models.User) (models.User, error) {
	user.ID = uuid.New().String()
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hash)

	if result := s.db.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *UserRepository) AuthenticateUser(login, pass string) (models.User, error) {
	var user models.User
	if result := s.db.Where("username = ?", login).Or("email = ?", login).First(&user); result.Error != nil {
		return models.User{}, models.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass)); err != nil {
		return models.User{}, models.ErrInvalidCredentials
	}

	return user, nil
}

func (s *UserRepository) GetUserByID(id string) (models.User, error) {
	var user models.User
	if result := s.db.First(&user, "id = ?", id); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (s *UserRepository) getUserByUsername(username string) (models.User, error) {
	var user models.User
	if result := s.db.Where("username = ?", username).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
