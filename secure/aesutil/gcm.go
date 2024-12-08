package aesutil

import (
	"crypto/cipher"
)

// @todo don't work
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

// @todo don't work
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
	if len(cipherText) < nonceSize {
		return nil, err
	}

	return gcm.Open(nil, nonce, cipherText, nil)
}
