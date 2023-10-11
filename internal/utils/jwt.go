package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Role  string `json:"role"`
	Email string `json:"email"`
	jwt.StandardClaims
}

var (
// Keys used to sign and verify our tokens
// verifyKey *rsa.PublicKey
// currently unused: signKey   *rsa.PrivateKey
// verifyKey = []byte("aaa")
)

var signingKey = []byte(os.Getenv("SECRET_KEY"))

// func ParseAndCheckToken(token string) (*Claims, error) {
// 	// the API key is a JWT signed by us with a claim to be a reseller
// 	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(parsedToken *jwt.Token) (interface{}, error) {
// 		// the key used to validate tokens
// 		return verifyKey, nil
// 	})

// 	if err == nil {
// 		if claims, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
// 			return claims, nil
// 		}
// 	}
// 	return nil, err
// }

func ParseAndCheckToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func CreateToken(email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
