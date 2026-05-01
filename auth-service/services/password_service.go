package services

import (
	"crypto/rand"
	"math/big"
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
