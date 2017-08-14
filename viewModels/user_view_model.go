package viewModels

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserViewModel struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	UserName  string        `form:"username" json:"username" binding:"required"`
	FirstName string        `form:"firstName" json:"firstName" binding:"required"`
	LastName  string        `form:"lastName" json:"lastName" binding:"required"`
	Slug      string        `json:"slug"`
	Password  string        `form:"password" json:"password"`
	Email     string        `form:"email" json:"email" binding:"required"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	UpdaterId string        `json:"updaterId"`
}

func (r *UserViewModel) IsValid() error {
	return nil
}
