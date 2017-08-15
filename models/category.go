package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Category struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	UpdaterId string        `json:"updaterId"`
}
