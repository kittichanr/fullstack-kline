package usecase

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/pkg"
	"errors"
)

type AuthUseCase interface {
	SetPin(name, pin string) error
	AuthenticatePin(name, pin string) (string, error)
}

type authUseCase struct {
	repo   repository.UserRepository
	secret string
}

func NewAuthUseCase(repo repository.UserRepository, secret string) AuthUseCase {
	return &authUseCase{repo, secret}
}

func (u *authUseCase) SetPin(name, pin string) error {
	if len(pin) != 6 {
		return errors.New("PIN must be 6 digits")
	}

	hashedPin, err := pkg.HashPin(pin)
	if err != nil {
		return err
	}

	user, err := u.repo.FindByName(name)
	if err != nil {
		return u.repo.CreateUser(&domain.User{Name: user.Name, PinHash: hashedPin})
	}

	return u.repo.UpdatePinHash(name, hashedPin)
}

func (u *authUseCase) AuthenticatePin(name, pin string) (string, error) {
	user, err := u.repo.FindByName(name)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !pkg.ComparePin(user.PinHash, pin) {
		return "", errors.New("invalid PIN")
	}

	token, err := pkg.GenerateToken(user.UserID, name, u.secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
