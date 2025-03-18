package gecko

import (
	cr "crypto/rand"
)

func GenerateSecureBytes(l int) ([]byte, error) {
	randBytes := make([]byte, l)

	_, err := cr.Read(randBytes)
	if err != nil {
		return nil, err
	}

	return randBytes, nil

}

// generates a random 32 byte key
func GenerateKey() ([]byte, error) {
	return GenerateSecureBytes(32)
}

// generate a random 16 byte IV
func GenerateIV() ([]byte, error) {
	return GenerateSecureBytes(16)
}
