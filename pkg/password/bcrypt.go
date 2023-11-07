package password

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword -.-
func HashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	return string(hashedPasswordBytes), err
}

// VerifyPassword -.-
func VerifyPassword(hashedPassword, currPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(currPassword)) == nil
}
