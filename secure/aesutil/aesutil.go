// Package aesutil provides aes encryption and decryption.
package aesutil

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const (
	Aes128 = 16
	Aes192 = 24
	Aes256 = 32
)

// AdjustKey 获取对自适应长度秘钥
func AdjustKey(key []byte) (cipher.Block, error) {
	vLen := len(key)
	var adjustKey []byte
	if vLen >= Aes256 {
		adjustKey = key[:Aes256]
	} else if vLen >= Aes192 {
		adjustKey = key[:Aes192]
	} else if vLen >= Aes128 {
		adjustKey = key[:Aes128]
	} else {
		return nil, errors.New("秘钥长度太短，请提供256,192，128长度的秘钥")
	}

	block, err := aes.NewCipher(adjustKey)
	return block, err
}
