package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTS struct {
	AccessToken  string
	RefreshToken string
}
type JwtClaim struct {
	UserId string `json:"userId"`
	jwt.StandardClaims
}

func CreateJWT(userId primitive.ObjectID) (*JWTS, error) {
	payload := userId.Hex()
	accessToken, accessTokenErr := accessToken(payload)
	if accessTokenErr != nil {
		return nil, accessTokenErr
	}
	refreshToken, refreshTokenErr := refreshToken(payload)

	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}
	return &JWTS{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func accessToken(UserId string) (string, error) {
	claims := &JwtClaim{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(30)).Unix(),
		},
	}
	ss, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("ACCESS_TOKEN_S")))
	if err != nil {
		log.Panic(err)
	}
	return ss, err
}
func refreshToken(UserId string) (string, error) {
	claims := &JwtClaim{
		UserId: UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1000)).Unix(),
		},
	}
	ss, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("REFRESH_TOKEN_S")))
	if err != nil {
		log.Panic(err)
	}
	return ss, err
}
