package repo

import (
	"HarlanCinema/pkg/models"

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

func (s *UserRepository) Create(user models.User) (models.User, error) {
	if err := s.db.Create(&user).Error; err != nil {
		return models.User{}, err
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

func (s *UserRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if result := s.db.Where("username = ?", username).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
