package password

import "golang.org/x/crypto/bcrypt"

// Encrypt returns a hash string
func Encrypt(password string) (encryptedString string, err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hash), nil
}

// Compare will return boolian and check if the password matches
func Compare(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
