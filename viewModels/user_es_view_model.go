package viewModels

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserEsViewModel struct {
	Id           bson.ObjectId `bson:"_id" json:"id"`
	UserName     string        `form:"username" json:"username" binding:"required"`
	UserType     string        `form:"userType" json:"userType"`
	FirstName    string        `form:"firstName" json:"firstName" binding:"required"`
	LastName     string        `form:"lastName" json:"lastName" binding:"required"`
	RealName     string        `form:"realName" json:"realName"`
	Image        string        `form:"image" json:"image"`
	Slug         string        `json:"slug"`
	Email        string        `form:"email" json:"email" binding:"required"`
	Title        string        `form:"title" json:"title"`
	AboutMe      string        `form:"aboutMe" json:"aboutMe"`
	Phone        string        `form:"phone" json:"phone"`
	CountryName  string        `form:"countryName" json:"countryName"`
	MinutePrice  float64       `form:"minutePrice" json:"minutePrice"`
	Rate         float64       `form:"rate" json:"rate"`
	ReviewsCount float64       `form:"reviewsCount" json:"reviewsCount"`
	BirthDay     time.Time     `bson:"birthday" form:"birthDay" json:"birthDay"`
	Tags         []string      `form:"tags" json:"tags"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
}
