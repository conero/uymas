package aesutil

import "bytes"

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padText...)
}

func ZeroUnPadding(origData []byte) []byte {
	for i := len(origData) - 1; ; i-- {
		if origData[i] != 0 {
			return origData[:i+1]
		}
	}
}
