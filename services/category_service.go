package services

import (
	"strings"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/enkhalifapro/pgen/db"
	"github.com/enkhalifapro/pgen/utilities"
	"github.com/enkhalifapro/pgen/models"
	"github.com/enkhalifapro/pgen/viewModels"
)

// categoryCollection default name.
var categoryCollection = "categories"

type CategoryService struct {
	DB          *db.DB              `inject:""`
	SlugUtil    *utilities.SlugUtil `inject:""`
	UserService *UserService        `inject:""`
}

func (r *CategoryService) QueryByPage(query *bson.M, pageSize int, pageNumber int) (categories []*models.Category, count int) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(categoryCollection)

	skip := (pageNumber - 1) * pageSize

	iterator := collection.Find(query).Limit(pageSize).Skip(skip).Iter()
	categories = make([]*models.Category, 0)
	category := models.Category{}
	for iterator.Next(&category) {
		newCategory := category
		categories = append(categories, &newCategory)
	}
	count, _ = collection.Find(query).Count()

	return categories, count
}

func (r *CategoryService) ResolveCategory(categoryId string) (category *viewModels.CategoryViewModel) {
	category = &viewModels.CategoryViewModel{}
	// validate authorId
	isValidId := bson.IsObjectIdHex(categoryId)
	if isValidId == false {
		return category
	}

	category, _ = r.FindById(categoryId)

	return category
}

func (r *CategoryService) FindById(id string) (*viewModels.CategoryViewModel, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(categoryCollection)
	category := viewModels.CategoryViewModel{}
	err := collection.FindId(bson.ObjectIdHex(id)).One(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// FindOne returns first instance matched to query.
func (r *CategoryService) FindOne(query *bson.M) (*viewModels.CategoryViewModel, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	var category viewModels.CategoryViewModel
	if err := session.DB("").C(categoryCollection).Find(query).One(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryService) FindAll(query *bson.M) ([]*models.Category, error) {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(categoryCollection)
	categories := make([]*models.Category, 0)
	err := collection.Find(query).All(&categories)
	return categories, err
}

func (r *CategoryService) Insert(updaterId string, category *models.Category) error {
	category.Name = strings.ToLower(category.Name)
	category.ID = bson.NewObjectId()
	category.CreatedAt = time.Now().UTC()
	category.UpdatedAt = time.Now().UTC()
	category.UpdaterId = updaterId
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(categoryCollection)
	err := collection.Insert(category)
	return err
}

func (r *CategoryService) Update(updaterId string, id string, newCategory *models.Category) error {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(categoryCollection)
	err := collection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{
		"name":      newCategory.Name,
		"updatedat": time.Now().UTC()}})
	return err
}

func (r *CategoryService) Delete(updaterId string, id string) error {
	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(categoryCollection)
	err := collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
