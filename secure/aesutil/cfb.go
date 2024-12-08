package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func CfbEncrypt(data, key, iv []byte) (ciphertext []byte, err error) {
	mode, er := AdjustKey(key)
	if er != nil {
		return nil, er
	}

	iv = checkAndGenIv(key, iv, mode.BlockSize())
	ciphertext = make([]byte, aes.BlockSize+len(data))

	stream := cipher.NewCFBEncrypter(mode, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)
	return
}

func CfbDecrypt(cipherText, key, iv []byte) (rawtext []byte, err error) {
	mode, er := AdjustKey(key)
	if er != nil {
		return nil, er
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv = checkAndGenIv(key, iv, mode.BlockSize())

	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(mode, iv)
	stream.XORKeyStream(cipherText, cipherText)

	rawtext = cipherText
	return
}
