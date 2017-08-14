package services

import (
	"math"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gitlab.com/enkhalifapro/ulearn-api/db"
	"gitlab.com/enkhalifapro/ulearn-api/utilities"
	"gitlab.com/enkhalifapro/ulearn-api/viewModels"
)

var (
	// UserCollection default name.
	UserCollection = "users"

	// userSessionExpirePeriod used when we create session object after login operation.
	userSessionExpirePeriod = time.Hour * 24 * 31
)

type UserService struct {
	DB               *db.DB               `inject:""`
	SessionService   *SessionService      `inject:""`
	CryptUtil        *utilities.CryptUtil `inject:""`
	MandrillMailUtil *utilities.MandrillMailUtil  `inject:""`
	SlugUtil         *utilities.SlugUtil  `inject:""`
}

func (r *UserService) QueryByPage(query *bson.M, pageSize int, pageNumber int) ([]*viewModels.UserViewModel, int, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(UserCollection)
	skip := (pageNumber - 1) * pageSize
	set := collection.Find(query)
	if pageSize > 0 {
		set = set.Limit(pageSize)
	}
	if skip > 0 {
		set = set.Skip(skip)
	}

	users := make([]*viewModels.UserViewModel, 0)
	user := viewModels.UserViewModel{}
	iterator := set.Iter()
	for iterator.Next(&user) {
		newUser := user
		if math.IsNaN(user.Rate) == true {
			newUser.Rate = 0
		}
		newUser.CreationDate = user.CreatedAt.UTC().Format(time.RFC3339)
		users = append(users, &newUser)
	}

	if pageSize != 0 {
		count, err := collection.Find(query).Count()
		if err != nil {
			return nil, 0, err
		}
		return users, count, nil
	}
	return users, len(users), nil
}
