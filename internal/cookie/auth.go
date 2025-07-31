package cookie

import (
	"errors"
	"fmt"
	"gotube/internal/user"
	"net/http"
	"strings"
)

func hasRoleWithKey(r *http.Request, role string, keyIndex int) bool {
	cookie, err := r.Cookie("auth")
	if err != nil {
		// Cookie not found, assume the user is not authenticated
		return false
	}

	value, err := verifySignedCookieWithKey("auth", cookie.Value)
	if err != nil {
		// Invalid cookie signature or format, assume the user is not authenticated
		return false
	}

	parts := strings.Split(value, "|")
	if len(parts) != 2 {
		fmt.Println("3")
		// Invalid cookie format, assume the user is not authenticated

		return false
	}

	username := parts[0]
	roleValue := parts[1]

	// Verify that the user has the required role
	for _, u := range user.Users {
		if u.Username == username && u.Role == roleValue {
			return u.Role == role
		}
	}

	// The user was not found, assume the user is not authenticated
	return false
}

// userRole examines the request’s cookies in priority order.
// It returns "admin" or "user" if a matching role cookie is found,
// or an error if no valid role is present.
func UserRole(r *http.Request) (string, error) {
	// Check for an admin first, then user
	for key := range cookieKeys {
		if hasRoleWithKey(r, "admin", key) {
			return "admin", nil
		}
		if hasRoleWithKey(r, "user", key) {
			return "user", nil
		}
	}
	return "", errors.New("no valid user role found in cookies")
}
