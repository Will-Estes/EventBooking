package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SECRET = "some-key"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email_address": email,
		"user_id":       userId,
		"exp":           time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(SECRET))
}

// Returns userId (or -1) and flag for valid token
func IsTokenValid(token string) (int64, bool) {
	if token == "" {
		return -1, false
	}

	token = token[len("Bearer "):]

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			fmt.Println("invalid signing method")
			return nil, errors.New("invalid signing method")
		}

		return []byte(SECRET), nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return -1, false
	}

	if !parsedToken.Valid {
		fmt.Println("invalid token")
		return -1, false
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("invalid token claims")
		return -1, false
	}

	//email := claims["email_address"].(string)
	userId := int64(claims["user_id"].(float64))

	return userId, true
}
