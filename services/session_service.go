package services

import (
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/enkhalifapro/pgen/db"
	"github.com/enkhalifapro/pgen/utilities"
	"github.com/enkhalifapro/pgen/models"
)

// sessionCollection default name.
var sessionCollection = "sessions"

// SessionService for auth users.
type SessionService struct {
	DB        *db.DB               `inject:""`
	CryptUtil *utilities.CryptUtil `inject:""`
}

// Insert session to database.
func (r *SessionService) Insert(userSession *models.Session) error {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	userSession.Id = bson.NewObjectId()
	userSession.Token = r.CryptUtil.Encrypt(r.CryptUtil.RandomString(100))
	return session.DB("").C(sessionCollection).Insert(userSession)
}

// Valid checks is session a live.
func (r *SessionService) Valid(sessionToken string) bool {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(sessionCollection)
	// TODO: add logging
	count, _ := collection.Find(&bson.M{
		"token": sessionToken,
		"expirydate": bson.M{
			"$gte": time.Now().UTC(),
		},
	}).Count()
	return count > 0
}

// Find session in database and return.
func (r *SessionService) Find(sessionToken string) (*models.Session, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	userSession := &models.Session{}
	err := session.DB("").C(sessionCollection).Find(&bson.M{
		"token": sessionToken,
		"expirydate": bson.M{
			"$gte": time.Now().UTC(),
		},
	}).One(userSession)
	if err != nil {
		return nil, err
	}
	return userSession, nil
}

// Logout makes session outdated.
func (r *SessionService) Logout(token string) error {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(sessionCollection)
	return collection.Update(
		bson.M{"token": token},
		bson.M{
			"$set": bson.M{
				"expirydate": time.Now().UTC(),
			},
		},
	)
}
