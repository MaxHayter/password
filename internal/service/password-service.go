package service

import (
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CheckAndHash(ctx context.Context, password string) ([]byte, error) {
	numSymbolChars := `!@$&*-., "#%'()+/:;<=>?[\]^{|}~`
	upperChars := `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	digitsChars := `0123456789`

	pwLength := len(password)
	numeric := 0
	numSymbols := 0
	upper := 0

	for _, c := range password {
		switch {
		case strings.ContainsRune(numSymbolChars, c):
			numSymbols++
		case strings.ContainsRune(digitsChars, c):
			numeric++
		case strings.ContainsRune(upperChars, c):
			upper++
		}
	}

	if upper > 3 {
		upper = 3
	}
	if numSymbols > 3 {
		numSymbols = 3
	}
	if numeric > 3 {
		numeric = 3
	}

	if (pwLength-20)+(numeric)+(numSymbols*2)+(upper) < 0 {
		return nil, errors.New("weak password")
	}

	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (s *Service) Compare(ctx context.Context, password string, hash []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil, err
}
