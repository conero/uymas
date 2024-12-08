package aesutil

import (
	"gitee.com/conero/uymas/v2/str"
	"testing"
)

func TestCfbEncrypt(t *testing.T) {
	key := []byte(str.RandStr.String(32))
	iv := []byte(str.RandStr.String(16))

	// 加密
	data := "test in aes/CFB. 路人借问遥招手，怕得鱼惊不应人。"
	encrypt, err := CfbEncrypt([]byte(data), key, iv)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("加密成功！密文：%s", string(encrypt))
	}

	// 解密
	decrypt, err := CfbDecrypt(encrypt, key, iv)
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
