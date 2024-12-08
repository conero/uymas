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
func CbcEncrypt(data, key, iv []byte) (ciphertext []byte, err error) {
	block, err := AdjustKey(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	if len(data)%blockSize != 0 {
		data = ZeroPadding(data, block.BlockSize())
	}

	ciphertext = make([]byte, len(data))
	iv = checkAndGenIv(key, iv, blockSize)

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, data)
	return ciphertext, nil
}

func CbcDecrypt(cipherText, key, iv []byte) (rawtext []byte, err error) {
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

	rawtext = make([]byte, len(cipherText))
	iv = checkAndGenIv(key, iv, blockSize)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(rawtext, cipherText)
	rawtext = ZeroUnPadding(rawtext)
	return rawtext, nil
}
