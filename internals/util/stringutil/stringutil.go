package stringutil

import (
	b64 "encoding/base64"
)

// HashFromString generate hash from string
func HashFromString(s string) (hash string) {
	sEnc := b64.StdEncoding.EncodeToString([]byte(s))
	return sEnc
}

// StringFromHash convert hash to string
func StringFromHash(s string) (str string, err error) {
	sDec, err := b64.StdEncoding.DecodeString(s)
	return string(sDec), err
}
