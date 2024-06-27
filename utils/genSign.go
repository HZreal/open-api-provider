package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenSign(params, secretKey string) string {
	// md5
	hash := md5.New()
	hash.Write([]byte(params + "." + secretKey))
	return hex.EncodeToString(hash.Sum(nil))
}
