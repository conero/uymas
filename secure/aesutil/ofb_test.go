package aesutil

import (
	"gitee.com/conero/uymas/v2/str"
	"testing"
)

// TestOfbEncrypt tests the OfbEncrypt function for proper encryption.
func TestOfbEncrypt(t *testing.T) {
	key := []byte(str.RandStr.String(32))
	iv := []byte(str.RandStr.String(16))

	// 加密
	data := "test in aes/Ofb.  墙角数枝梅，凌寒独自开；遥知不是雪，为有暗香来。~"
	encrypt, err := OfbEncrypt([]byte(data), key, iv)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("加密成功！密文：%s", string(encrypt))
	}

	// 解密
	decrypt, err := OfbDecrypt(encrypt, key, iv)
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
