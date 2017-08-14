package services

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/enkhalifapro/pgen/db"
	"github.com/enkhalifapro/pgen/utilities"
	"github.com/enkhalifapro/pgen/viewModels"
	"github.com/enkhalifapro/pgen/models"
)

var (
	// ServiceCollection default name.
	ServiceCollection = "services"
)

type ServiceService struct {
	DB *db.DB               `inject:""`
}

func (r *UserService) FindByName(name string) (*models.Service, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(ServiceCollection)
	service := models.Service{}
	err := collection.Find(bson.M{"name": name}).One(&service)
	if err != nil {
		return nil, err
	}
	return &service, nil
}
