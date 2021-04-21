package crypto

import "golang.org/x/crypto/bcrypt"

func Ecrypting(texto string) (string, error) {
	cons := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(texto), cons)
	return string(bytes), err
}
