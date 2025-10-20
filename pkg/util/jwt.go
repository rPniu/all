package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rPniu/all/pkg/setting"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Uid string `json:"uid"`
	jwt.RegisteredClaims
}

func GenerateToken(Uid string) (string, error) {
	claims := Claims{
		Uid: Uid,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
func ParseToken(tokenstr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenstr,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			// 验证签名方法是否为预期的 HS256（防止算法被篡改）
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			// 返回用于验证的密钥
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
