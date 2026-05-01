package services

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const temporaryPasswordLength = 12

const lowercaseChars = "abcdefghijklmnopqrstuvwxyz"
const uppercaseChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitChars = "0123456789"
const specialChars = "!@#$%^&*()-_=+[]{}<>?"

var passwordCharset = []string{lowercaseChars, uppercaseChars, digitChars, specialChars}

func GenerateTemporaryPassword() string {
	password := make([]byte, 0, temporaryPasswordLength)
	for _, charset := range passwordCharset {
		char, err := randomCharFrom(charset)
		if err != nil {
			return "TempPass123!"
		}
		password = append(password, char)
	}

	allChars := lowercaseChars + uppercaseChars + digitChars + specialChars
	for len(password) < temporaryPasswordLength {
		char, err := randomCharFrom(allChars)
		if err != nil {
			return "TempPass123!"
		}
		password = append(password, char)
	}

	if err := shuffleBytes(password); err != nil {
		return "TempPass123!"
	}

	return string(password)
}

func randomCharFrom(charset string) (byte, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		return 0, err
	}

	return charset[index.Int64()], nil
}

func shuffleBytes(data []byte) error {
	for i := len(data) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}

		swapIndex := int(j.Int64())
		data[i], data[swapIndex] = data[swapIndex], data[i]
	}

	return nil
}

func ValidatePasswordStrength(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	if strings.ContainsAny(password, " \t\n\r") {
		return errors.New("password must not contain spaces")
	}
	if !containsCharFromSet(password, lowercaseChars) {
		return errors.New("password must contain at least one lowercase letter")
	}
	if !containsCharFromSet(password, uppercaseChars) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !containsCharFromSet(password, digitChars) {
		return errors.New("password must contain at least one number")
	}
	if !containsCharFromSet(password, specialChars) {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func containsCharFromSet(value string, charset string) bool {
	for _, char := range value {
		if strings.ContainsRune(charset, char) {
			return true
		}
	}

	return false
}
