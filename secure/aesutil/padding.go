package aesutil

import (
	"bytes"
	"errors"
)

func ZeroPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(data, padText...)
}

func ZeroUnPadding(data []byte) []byte {
	for i := len(data) - 1; ; i-- {
		if data[i] != 0 {
			return data[:i+1]
		}
	}
}

func Pkcs7Padding(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, errors.New("block size must be greater than 0")
	}

	padding := blockSize - (len(data) % blockSize)
	paddedData := make([]byte, len(data)+padding)
	copy(paddedData, data)
	for i := len(data); i < len(paddedData); i++ {
		paddedData[i] = byte(padding)
	}
	return paddedData, nil
}

func Pkcs7UnPadding(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid padded data length")
	}

	padding := int(data[len(data)-1])

	if padding > blockSize || padding == 0 {
		return nil, errors.New("invalid padding value")
	}

	for i := len(data) - padding; i < len(data); i++ {
		if data[i] != byte(padding) {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:len(data)-padding], nil
}
