package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ilyasssan2/golangGin/entity"
	"github.com/ilyasssan2/golangGin/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	collection := utils.Client.Database(os.Getenv("DB_NAME")).Collection("user")
	var body loginReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	var user entity.User
	if err := collection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry there is no user with this credentials"})
		return
	}
	if isPasswordValid := utils.ComparePassword(body.Password, user.Password); !isPasswordValid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "sorry there is no user with this credentials"})
		return
	}
	jwts, err := utils.CreateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.SetCookie("jid", jwts.RefreshToken, int((time.Hour * time.Duration(1000)).Seconds()), "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"accessToken": jwts.AccessToken})
}

type registerReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	collection := utils.Client.Database(os.Getenv("DB_NAME")).Collection("user")
	var body registerReq
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	searchingCtx, cancelSearching := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelSearching()
	if collection.FindOne(searchingCtx, bson.M{"email": body.Email}).Err() == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exist please try to login"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	hashedPassword, err := utils.HashPasword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	body.Password = hashedPassword
	res, err := collection.InsertOne(ctx, body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var jwts *utils.JWTS
	jwts, err = utils.CreateJWT(res.InsertedID.(primitive.ObjectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.SetCookie("jid", jwts.RefreshToken, int((time.Hour * time.Duration(1000)).Seconds()), "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{
		"accessToken": jwts.AccessToken})
}
