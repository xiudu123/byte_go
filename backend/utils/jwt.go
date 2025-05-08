package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

/**
 * @author: 锈渎
 * @date: 2025/2/11 21:58
 * @code: 面向对象面向君， 不负代码不负卿。
 * @description: jwt 工具类

github.com/golang-jwt/jwt/v5
*/

var SecretKey = []byte("sgfhgjkyukjsedfsdijfdxjknvbbxcklgmdsifjskljfnaszkldfgdfhfgfnsdkmgnsdlkgf")

type Claims struct {
	UserId uint32 `json:"user_id"`
	//Role   string `json:"role"`
	JTI string `json:"jti"`
	jwt.RegisteredClaims
}

var TokenTTL = 24 * time.Hour

func GenerateToken(userId uint32, permissionVersion int64) (string, string, error) {
	secretKey := append(SecretKey, []byte(fmt.Sprintf("%d", permissionVersion))...)
	jti := fmt.Sprintf("jwt-%d-%d", userId, time.Now().Unix())
	claims := &Claims{
		UserId: userId,
		//Role:   role,
		JTI: jti,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenTTL)),
			Issuer:    "byte_go",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	return tokenString, jti, err
}

func ParseToken(tokenString string, permissionVersion int64) (*Claims, error) {
	secretKey := append(SecretKey, []byte(fmt.Sprintf("%d", permissionVersion))...)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
