package viewModels

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserCartViewModel struct {
	Id        bson.ObjectId    `bson:"_id" json:"id"`
	UserId    string           `json:"userId"`
	CourseId  string           `json:"courseId"`
	Course    *CourseViewModel `json:"course"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	UpdaterId string           `json:"updaterId"`
}
