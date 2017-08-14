package viewModels

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UserCourseProgressViewModel struct {
	Id         bson.ObjectId    `bson:"_id" json:"id"`
	UserId     string           `json:"userId"`
	CourseId   string           `json:"courseId" binding:"required"`
	Course     *CourseViewModel `json:"course"`
	Type       string           `json:"type" binding:"required"`
	LessonId   string           `json:"lessonId" binding:"required"`
	LessonSlug string           `json:"lessonSlug" binding:"required"`
	PageId     string           `json:"pageId" binding:"required"`
	PageSlug   string           `json:"pageSlug" binding:"required"`
	CreatedAt  time.Time        `json:"createdAt"`
	UpdatedAt  time.Time        `json:"updatedAt"`
	UpdaterId  string           `json:"updaterId"`
}

func (r UserCourseProgressViewModel) IsValid() error {
	if bson.IsObjectIdHex(r.UserId) == false {
		return errors.New("Invalid userId")
	}
	if bson.IsObjectIdHex(r.CourseId) == false {
		return errors.New("Invalid courseId")
	}
	if bson.IsObjectIdHex(r.LessonId) == false {
		return errors.New("Invalid lessonId")
	}
	if bson.IsObjectIdHex(r.PageId) == false {
		return errors.New("Invalid pageId")
	}
	return nil
}
