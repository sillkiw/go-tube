package cookie

import (
	"crypto/hmac"
	"crypto/sha256"
)

func computeCookieSignature(name, value string, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(name))
	h.Write([]byte("|"))
	h.Write([]byte(value))
	return h.Sum(nil)
}

func validateCookieSignature(name, value string, signature []byte) bool {
	expectedSignature := computeCookieSignature(name, value, cookieKeys[currentKeyIndex])
	return hmac.Equal(signature, expectedSignature)
}
