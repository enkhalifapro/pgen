package viewModels

import "gopkg.in/mgo.v2/bson"

type CategoryViewModel struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string `json:"name"`
}
