package viewModels

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserMessageViewModel struct {
	Id         bson.ObjectId  `bson:"_id" json:"id"`
	UserId     string         `form:"userId" json:"userId" binding:"required"`
	User       *UserViewModel `bson:"-" json:"user"`
	ProducerId string         `form:"producerId" json:"producerId" binding:"required"`
	Producer   *UserViewModel `bson:"-" json:"producer"`
	Title      string         `form:"title" json:"title" binding:"required"`
	Content    string         `form:"content" json:"content" binding:"required"`
	Link       string         `form:"link" json:"link" binding:"required"`
	Seen       bool           `form:"seen" json:"seen"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	UpdaterId  string         `json:"updaterId"`
}
