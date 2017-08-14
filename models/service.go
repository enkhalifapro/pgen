package models

import "gopkg.in/mgo.v2/bson"

type ServicePolicy struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string `json:"name"`
	Categories []string `json:"categories"`
	Policy     string `json:"policy"`
}
