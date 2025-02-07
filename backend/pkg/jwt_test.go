package pkg_test

import (
	"backend/pkg"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	userID := "123"
	name := "John Doe"
	jwtSecret := "secret"

	token, err := pkg.GenerateToken(userID, name, jwtSecret)
	if err != nil {
		t.Fatalf("GenerateToken returned an error: %v", err)
	}

	if len(token) == 0 {
		t.Fatal("Generated token should not be empty")
	}
}

func TestValidateToken(t *testing.T) {
	userID := "123"
	name := "John Doe"
	jwtSecret := "secret"

	token, err := pkg.GenerateToken(userID, name, jwtSecret)
	if err != nil {
		t.Fatalf("GenerateToken returned an error: %v", err)
	}

	parsedToken, err := pkg.ValidateToken(token, jwtSecret)
	if err != nil {
		t.Fatalf("ValidateToken returned an error: %v", err)
	}

	if !parsedToken.Valid {
		t.Error("Parsed token should be valid")
	}
}

func TestValidateTokenWithInvalidSecret(t *testing.T) {
	userID := "123"
	name := "John Doe"
	jwtSecret := "secret"

	token, err := pkg.GenerateToken(userID, name, jwtSecret)
	if err != nil {
		t.Fatalf("GenerateToken returned an error: %v", err)
	}

	_, err = pkg.ValidateToken(token, "wrong_secret")
	if err == nil {
		t.Error("ValidateToken should return an error for an incorrect secret")
	}
}
