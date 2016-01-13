package utils

func EncryptPassword(password string) string {
	digest, err := GenerateHash(password)
	if err != nil {
		return ""
	}

	return string(digest)
}
