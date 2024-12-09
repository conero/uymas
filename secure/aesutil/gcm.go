package aesutil

import (
	"crypto/cipher"
	"errors"
)

func GcmEncrypt(data, key, nonce []byte) (ciphertext []byte, err error) {
	block, err := AdjustKey(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce = checkAndGenIv(key, nonce, gcm.NonceSize())
	ciphertext = gcm.Seal(nil, nonce, data, nil)
	return
}

func GcmDecrypt(cipherText, key, nonce []byte) (rawtext []byte, err error) {
	block, err := AdjustKey(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonceLen := len(nonce)
	if nonceLen < nonceSize {
		return nil, errors.New("nonce is too short")
	} else if nonceLen > nonceSize {
		nonce = nonce[:nonceSize]
	}

	return gcm.Open(nil, nonce, cipherText, nil)
}
