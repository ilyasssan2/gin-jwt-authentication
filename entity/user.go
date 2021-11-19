package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
}
