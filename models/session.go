package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Session struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	Token      string        `form:"userName" json:"userName"`
	ExpiryDate time.Time     `json:"expiryDate"`
	UserId     string        `json:"userId"`
	User       *User
	CreatedAt  bson.MongoTimestamp `json:"createdAt"`
}
