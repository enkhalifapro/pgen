package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Service struct {
	ID     bson.ObjectId `bson:"_id" json:"id"`
	Name   string `json:"name"`
	Policy string `json:"policy"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	UpdaterId string        `json:"updaterId"`
}
