package utils
import "golang.org/x/crypto/bcrypt"


func GenerateHash(password string) ([]byte, error)  {
	hex := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)
	if err != nil {
		return hashedPassword, err
	}

	return hashedPassword, nil
}
