package services

import (
	"dbo-backend/database"
	"dbo-backend/models"
	"fmt"
	"strconv"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type myClaims struct {
	jwt.StandardClaims
	Id    uint
	Email string
}

func GenerateToken(user models.User) (map[string]string, error) {
	acDuration, _ := strconv.ParseInt("3600", 10, 64)
	accessToken := models.AccessToken{
		UserID:    int(user.ID),
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add((time.Second * time.Duration(acDuration))),
		IsActive:  true,
	}
	database.Db.Create(&accessToken)

	newAccessToken, err := token(int32(user.ID), int32(accessToken.ID), acDuration)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token": newAccessToken,
	}, nil
}

// &jwt.NumericDate{ time.Now().Unix()}
func token(ID int32, TokenID int32, duration int64) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Subject:   fmt.Sprint(ID),
			ID:        fmt.Sprint(TokenID),
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add((time.Second * time.Duration(duration)))},
			Issuer:    "DBO",
		},
	)

	signedToken, err := token.SignedString([]byte("test123"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
