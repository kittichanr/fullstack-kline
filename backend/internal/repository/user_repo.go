package repository

import (
	"backend/internal/domain"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByName(name string) (*domain.User, error)
	CreateUser(user *domain.User) error
	UpdatePinHash(name, pinHash string) error
	GetGreetingByUserID(userID string) (*domain.UserGreeting, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByName(name string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) UpdatePinHash(name, pinHash string) error {
	return r.db.Model(&domain.User{}).Where("name = ?", name).Update("pin_hash", pinHash).Error
}

func (r *userRepository) GetGreetingByUserID(userID string) (*domain.UserGreeting, error) {
	var greeting domain.UserGreeting
	if err := r.db.Where("user_id = ?", userID).First(&greeting).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &greeting, nil
}
