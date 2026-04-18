package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("my-secret-key")

type Claims struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // 包含过期时间、签发者等标准字段
}

// 生成token
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(30 * time.Second) //30秒过期，测试用
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   "user_auth",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// 解析token
func ValidateJWT(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err // 可能是过期、签名错误、格式错误等
	}

	if !token.Valid {
		return nil, err //错误toekn
	}

	return claims, nil
}
