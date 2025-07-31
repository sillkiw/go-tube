package cookie

import "crypto/rand"

var cookieKeys [][]byte // Array of secret keys for key rotation
var currentKeyIndex int // Index of the current secret key

func InitializeKeys(n int) error {
	cookieKeys = nil
	for i := 0; i < n; i++ {
		key := make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			return err
		}
		cookieKeys = append(cookieKeys, key)
	}
	currentKeyIndex = 0
	return nil
}
