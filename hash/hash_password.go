package hash

import (
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

type passwordSha512 struct{}

func NewPasswordSha512() *passwordSha512 {
	return &passwordSha512{}
}
func (p *passwordSha512) GeneratePasswordSha512(passwordInput string) (string, error) {
	hash := sha512.New()
	_, err := hash.Write([]byte(passwordInput))
	if err != nil {
		return "", err
	}
	data := hex.EncodeToString(hash.Sum(nil))
	return data, nil
}

func (p *passwordSha512) ComparePassword(passInput, passHash string) (bool, error) {
	data, err := p.GeneratePasswordSha512(passInput)
	if err != nil {
		return false, err
	}
	if data != passHash {
		return false, errors.New("password does not match")
	}
	return true, nil
}
