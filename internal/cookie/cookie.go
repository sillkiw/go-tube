package cookie

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
	"time"
)

// CreateSignedCookie - a function that create a cookie.
func CreateSignedCookie(name, value string, expires time.Time) *http.Cookie {
	// Select the current cookie key (keys.go)
	cookieKey := cookieKeys[currentKeyIndex]
	// Encode the cookie value and sign it with the cookie key
	cookieValue := value + "|" + expires.Format(time.RFC3339)
	signature := computeCookieSignature(name, cookieValue, cookieKey) // signing.go

	cookieValueBase64 := base64.StdEncoding.EncodeToString([]byte(cookieValue))
	signatureBase64 := base64.StdEncoding.EncodeToString(signature)

	// Create the cookie with the encoded value and signature
	cookie := &http.Cookie{
		Name:     name,
		Value:    cookieValueBase64 + "|" + signatureBase64,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  expires,
	}
	return cookie
}

func verifySignedCookieWithKey(name, value string) (string, error) {
	// Decode the cookie value and signature from base64
	parts := strings.Split(value, "|")
	if value == "" || len(parts) != 2 {
		return "", errors.New("INVALID COOKIE FORMAT")
	}
	cookieValueBase64 := parts[0]
	signatureBase64 := parts[1]
	cookieValue, err := base64.StdEncoding.DecodeString(cookieValueBase64)
	if err != nil {
		return "", errors.New("INVALID COOKIE FORMAT")
	}

	signature, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return "", errors.New("INVALID COOKIE FORMAT")
	}

	// Verify the signature of the cookie using the specified key
	if !validateCookieSignature(name, string(cookieValue), signature) {
		return "", errors.New("INVALID COOKIE FORMAT")
	}

	// Check if the cookie has expired
	parts = strings.Split(string(cookieValue), "|")
	if len(parts) != 3 {
		return "", errors.New("INVALID COOKIE FORMAT")
	}

	expiration, err := time.Parse(time.RFC3339, string(string(parts[2])))
	if err != nil {
		return "", errors.New("INVALID COOKIE FORMAT")
	}
	if time.Now().After(expiration) {
		return "", errors.New("COOKIE HAS EXPIRED")
	}
	// Return the cookie value
	return string(parts[0] + "|" + parts[1]), nil
}
