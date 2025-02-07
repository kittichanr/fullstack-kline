package pkg

import "golang.org/x/crypto/bcrypt"

func HashPin(pin string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	return string(hashed), err
}

func ComparePin(hashedPin, pin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPin), []byte(pin))
	return err == nil
}
