package aesutil

import (
	"gitee.com/conero/uymas/v2/str"
	"testing"
)

func TestCtrEncrypt(t *testing.T) {
	key := []byte(str.RandStr.String(32))
	iv := []byte(str.RandStr.String(16))

	// 加密
	data := "test in aes/Gcm.  王子曰：仲永之通悟，受之天也。其受之天也，贤于材人远矣。卒之为众人，则其受于人者不至也。彼其受之天也，如此其贤也，不受之人，且为众人；今夫不受之天，固众人，又不受之人，得为众人而已耶？ "
	encrypt, err := CtrEncrypt([]byte(data), key, iv)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("加密成功！密文：%s", string(encrypt))
	}

	// 解密
	decrypt, err := CtrDecrypt(encrypt, key, iv)
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
