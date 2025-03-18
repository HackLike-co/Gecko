package gecko

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rc4"
)

// Use AES-GCM to encrypt a byte slice
// returns the encrypted bytes slice and any errors
func AES_CBCEncrypt(payload []byte, key []byte, iv []byte) ([]byte, error) {
	bPayload := pad(payload)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ePayload := make([]byte, len(bPayload))

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ePayload, bPayload)

	return ePayload, nil
}

func pad(payload []byte) []byte {
	padding := (aes.BlockSize - len(payload)%aes.BlockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(payload, padText...)
}

func RC4_Encrypt(payload []byte, key []byte) ([]byte, error) {
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ePayload := make([]byte, len(payload))

	cipher.XORKeyStream(ePayload, payload)

	return ePayload, nil
}
