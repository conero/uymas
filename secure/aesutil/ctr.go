package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
)

func CtrEncrypt(plaintext, key, iv []byte) (ciphertext []byte, err error) {
	block, er := AdjustKey(key)
	if er != nil {
		return nil, er
	}

	iv = checkAndGenIv(key, iv, block.BlockSize())
	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return
}

func CtrDecrypt(cipherText, key, iv []byte) (plaintext []byte, err error) {
	block, er := AdjustKey(key)
	if er != nil {
		return nil, er
	}

	iv = checkAndGenIv(key, iv, block.BlockSize())
	plaintext = make([]byte, len(cipherText))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, cipherText[aes.BlockSize:])
	return
}
