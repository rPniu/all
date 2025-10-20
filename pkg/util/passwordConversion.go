package util

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GetHashBytes(password string) ([]byte, error) {
	if len(password) > 72 {
		return nil, errors.New("密码长度不能超过72个字符")
	}
	rawPassword := []byte(password)
	hashBytes, err := bcrypt.GenerateFromPassword(rawPassword, 10)
	if err != nil {
		return []byte{}, err
	}
	return hashBytes, nil
}

func GetPasswordHash(password string) (string, error) {
	hashBytes, err := GetHashBytes(password)
	if err != nil {
		return "", err
	}
	hashStr := string(hashBytes)
	return hashStr, nil
}

func ComparePasswordHashAndPassword(passwordHash, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
