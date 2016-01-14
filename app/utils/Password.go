package utils

func EncryptPassword(password string) string {
	digest, err := GenerateHash(password)
	if err != nil {
		return ""
	}

	return string(digest)
}

func ComparePassword(password, dbPassword string) bool {
	hex := []byte(dbPassword)
	return CompareHash(hex, password)
}
