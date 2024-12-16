// Package secure the tool include hash algorithm
package secure

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"gitee.com/conero/uymas/v2/rock"
	"hash"
	"strings"
)

func HashHmac(h func() hash.Hash, data, key []byte) []byte {
	vHash := hmac.New(h, key)
	vHash.Write(data)
	return vHash.Sum(nil)
}

// HashHmacString simplified mach hash string generation for fast digest generation, support sha224, sha256, sha384, sha512, md5
func HashHmacString(data, key string, algos ...string) string {
	algo := rock.Param("sha256", algos...)
	algo = strings.ToLower(algo)
	switch algo {
	case "sha224":
		return fmt.Sprintf("%x", HashHmac(sha256.New224, []byte(data), []byte(key)))
	case "sha256":
		return fmt.Sprintf("%x", HashHmac(sha256.New, []byte(data), []byte(key)))
	case "sha384":
		return fmt.Sprintf("%x", HashHmac(sha512.New384, []byte(data), []byte(key)))
	case "sha512":
		return fmt.Sprintf("%x", HashHmac(sha512.New, []byte(data), []byte(key)))
	case "md5":
		return fmt.Sprintf("%x", HashHmac(md5.New, []byte(data), []byte(key)))
	}
	return ""
}

func Hash(h func() hash.Hash, data []byte) []byte {
	vHash := h()
	vHash.Write(data)
	return vHash.Sum(nil)
}

// HashString simplified hash string generation for fast digest generation, support sha224, sha256, sha384, sha512, md5
func HashString(data string, algos ...string) string {
	algo := rock.Param("sha256", algos...)
	algo = strings.ToLower(algo)
	switch algo {
	case "sha224":
		return fmt.Sprintf("%x", Hash(sha256.New224, []byte(data)))
	case "sha256":
		return fmt.Sprintf("%x", Hash(sha256.New, []byte(data)))
	case "sha384":
		return fmt.Sprintf("%x", Hash(sha512.New384, []byte(data)))
	case "sha512":
		return fmt.Sprintf("%x", Hash(sha512.New, []byte(data)))
	case "md5":
		return fmt.Sprintf("%x", Hash(md5.New, []byte(data)))
	}
	return ""
}
