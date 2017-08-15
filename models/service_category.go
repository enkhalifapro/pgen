package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type ServiceCategory struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	ServiceID  string `json:"serviceId"`
	CategoryID string `json:"categoryId"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	UpdaterId string        `json:"updaterId"`
}
