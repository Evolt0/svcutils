package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("qwetyuasdhjkab")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 产生token的函数
func GenerateToken(username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "svcutils",
			IssuedAt:  nowTime.Unix(),
		},
	}
	//
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

// ParseToken 验证token的函数
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	//
	return nil, err
}
