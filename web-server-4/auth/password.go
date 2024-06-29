package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidatePassword(password string, hashedPassword string) bool {
	hash := []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword([]byte(password), hash)
	return err == nil
}
