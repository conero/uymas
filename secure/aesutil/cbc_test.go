package aesutil

import (
	"crypto/aes"
	"encoding/base64"
	"gitee.com/conero/uymas/v2/str"
	"testing"
)

func TestCbcEncrypt(t *testing.T) {
	key := []byte(str.RandStr.String(32))
	iv := []byte(str.RandStr.String(16))

	// 加密
	data := "I'am Joshua Conero，欧耶>>| 古道西风瘦马，夕阳西下，断肠人在天涯。"
	encrypt, err := CbcEncrypt([]byte(data), key, iv)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("加密成功！密文：%s", string(encrypt))
	}

	// 解密
	decrypt, err := CbcDecrypt(encrypt, key, iv)
	decryptStr := string(decrypt)
	if err != nil {
		t.Error(err)
	} else if decryptStr != data {
		t.Logf("len Vs: %d VS %d", len(decryptStr), len(data))
		t.Errorf("加解密不匹配，解密后明文：\n%#v\n原始明文：\n%#v", decryptStr, data)
	} else {
		t.Logf("解密成功！明文：%s", decryptStr)
	}
}

// 与 php 互通
//
// 见文件： doc/test/php-insecure-aes/aes_test_fns.php
func TestCbcEncrypt_php(t *testing.T) {
	data := `Demo|十年生死两茫茫，不思量，自难忘|2018-07-01|㎡`
	key := "Y0^n4Dh47:dd7wOXyWJ9,jN-tv8jxY8i"
	iv := "[K$jCYBej-vLVDQY"

	refCipher := `E/6l1oVnotjo0gbovDloCuIa/CXdMysaOqQIWfZiYEpaJa+FjgFbESeVnZWoCeC7gNXsP/3Vd2OTB7X1b9jxE6K7O8bV2WCyv7ruSa8bbvM=`
	dataPkcs7, err := Pkcs7Padding([]byte(data), aes.BlockSize)
	if err != nil {
		t.Errorf("Pkcs7Padding 失败，%v", err)
	}
	by, err := CbcEncrypt(dataPkcs7, []byte(key), []byte(iv))
	if err != nil {
		t.Errorf("加密失败，%v", err)
		return
	}
	relCipher := base64.StdEncoding.EncodeToString(by)
	t.Logf("relCipher: %#v", relCipher)
	if refCipher != relCipher {
		//t.Errorf("加密结果不匹配，参考值：\n%s\n实际值：\n%s", refCipher, relCipher)
	}

	// 解密
	cipherBytes, err := base64.StdEncoding.DecodeString(refCipher)
	if err != nil {
		t.Errorf("php操作密文无法解析为字节，%v", err)
		return
	}
	rawTxt, err := CbcDecrypt(cipherBytes, []byte(key), []byte(iv))
	if err != nil {
		t.Errorf("php密文解密错误，%v", err)
		return
	}
	rawTxt, err = Pkcs7UnPadding(rawTxt, aes.BlockSize)
	if string(rawTxt) != data {
		t.Errorf("php密文解密结果不匹配，参考值：\n%#v\n实际值：\n%#v", data, string(rawTxt))
		return
	}
	t.Logf("解密结果：%s", string(rawTxt))

}
