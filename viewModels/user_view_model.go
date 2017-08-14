package viewModels

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserViewModel struct {
	Id                  bson.ObjectId `bson:"_id" json:"id"`
	UserName            string        `form:"username" json:"username"`
	FirstName           string        `form:"firstName" json:"firstName"`
	LastName            string        `form:"lastName" json:"lastName"`
	Title               string        `form:"title" json:"title"`
	AboutMe             string        `form:"aboutMe" json:"aboutMe"`
	Image               string        `form:"image" json:"image"`
	Slug                string        `json:"slug"`
	Email               string        `form:"email" json:"email"`
	Phone               string        `form:"phone" json:"phone"`
	MinutePrice         float64       `form:"minutePrice" json:"minutePrice"`
	Rate                float64       `form:"rate" json:"rate"`
	CallsCount          float64       `form:"callsCount" json:"callsCount"`
	ReviewsCount        float64       `form:"reviewsCount" json:"reviewsCount"`
	Tags                string        `form:"tags" json:"tags"`
	CountryName         string        `form:"countryName" json:"countryName"`
	FaceBookProfileLink string        `form:"faceBookProfileLink" json:"faceBookProfileLink"`
	TwitterProfileLink  string        `form:"twitterProfileLink" json:"twitterProfileLink"`
	LinkedInProfileLink string        `form:"linkedInProfileLink" json:"linkedInProfileLink"`
	ProviderType        string        `form:"providerType" json:"providerType"`
	ProviderId          string        `form:"providerId" json:"providerId"`
	CreatedAt           time.Time     `json:"-"`
	CreationDate        string        `json:"createdAt"`
}
