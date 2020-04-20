package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) ([]byte, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash, err
}
func ComparatePassword(hash, contraseña string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(contraseña))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return true
	}

}
