package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type UserUseCase interface {
	GetGreetingByUserID(userID string) (*domain.UserGreeting, error)
}

type userUseCase struct {
	repo   repository.UserRepository
	secret string
}

func NewUserUseCase(repo repository.UserRepository, secret string) UserUseCase {
	return &userUseCase{repo, secret}
}

func (u *userUseCase) GetGreetingByUserID(userID string) (*domain.UserGreeting, error) {
	greeting, err := u.repo.GetGreetingByUserID(userID)
	if err != nil {
		return nil, err
	}
	if greeting == nil {
		return nil, nil
	}
	return &domain.UserGreeting{
		UserID:    greeting.UserID,
		Greeting:  greeting.Greeting,
		DummyCol2: greeting.DummyCol2,
	}, nil
}
