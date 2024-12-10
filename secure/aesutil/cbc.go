package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func checkAndGenIv(key, iv []byte, blockSize int) []byte {
	if len(iv) == 0 {
		iv = key[:blockSize]
	} else if len(iv) >= blockSize {
		iv = iv[:blockSize]
	} else {
		iv = ZeroPadding(iv, blockSize)
	}
	return iv
}

// CbcEncrypt aes cbc encrypt, refer to https://pkg.go.dev/crypto/cipher@latest#example-NewCBCDecrypter
func CbcEncrypt(plaintext, key, iv []byte) (ciphertext []byte, err error) {
	block, err := AdjustKey(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(plaintext)%blockSize != 0 {
		plaintext = ZeroPadding(plaintext, block.BlockSize())
	}

	ciphertext = make([]byte, len(plaintext))
	iv = checkAndGenIv(key, iv, blockSize)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return ciphertext, nil
}

func CbcDecrypt(cipherText, key, iv []byte) (plaintext []byte, err error) {
	block, err := AdjustKey(key)
	if err != nil {
		return nil, err
	}
	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	blockSize := block.BlockSize()
	if len(cipherText)%blockSize != 0 {
		return nil, errors.New("ciphertext is not a multiple of the block size")
	}

	plaintext = make([]byte, len(cipherText))
	iv = checkAndGenIv(key, iv, blockSize)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, cipherText)
	plaintext = ZeroUnPadding(plaintext)
	return plaintext, nil
}
