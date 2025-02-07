package pkg_test

import (
	"backend/pkg"
	"testing"
)

func TestHashPin(t *testing.T) {
	pin := "123456"
	hashedPin, err := pkg.HashPin(pin)
	if err != nil {
		t.Fatalf("HashPin returned an error: %v", err)
	}

	if len(hashedPin) == 0 {
		t.Fatal("Hashed pin should not be empty")
	}
}

func TestComparePin(t *testing.T) {
	pin := "123456"
	hashedPin, err := pkg.HashPin(pin)
	if err != nil {
		t.Fatalf("HashPin returned an error: %v", err)
	}

	if !pkg.ComparePin(hashedPin, pin) {
		t.Error("ComparePin should return true for a correct pin")
	}

	if pkg.ComparePin(hashedPin, "1234567") {
		t.Error("ComparePin should return false for an incorrect pin")
	}
}
