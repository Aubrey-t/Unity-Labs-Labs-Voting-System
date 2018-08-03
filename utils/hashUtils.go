package utils

import (
	"crypto/rand"
	//"crypto/sha1"
	//"encoding/base64"
	"fmt"
	"io"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/scrypt"
)

const pwHashBytes = 64

func init() {}

func GenerateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func GeneratePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h), nil
}

func GenerateUUID() (string, error) {
	uu , err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return uu.String(), nil
}