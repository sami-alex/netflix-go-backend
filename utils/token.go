package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const secretKey = "mysecretkey"

func GenerateToken(userId primitive.ObjectID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": userId.Hex(), // Convert ObjectID to hex string
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (primitive.ObjectID, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("couldn't parse token: %v", err)
	}

	if !parsedToken.Valid {
		return primitive.NilObjectID, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return primitive.NilObjectID, errors.New("invalid token claims")
	}

	userIDHex, ok := claims["_id"].(string)
	if !ok {
		return primitive.NilObjectID, errors.New("missing or invalid user ID in token claims")
	}

	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to convert hex to ObjectID: %v", err)
	}
	return userID, nil
}
