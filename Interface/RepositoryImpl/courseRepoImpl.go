package repository

import (
	"Basic-Edu-Program/Domain/model"
	"Basic-Edu-Program/UseCase/repository"

	"gorm.io/gorm"
)


type courseRepo struct {
	db *gorm.DB
}

func NewCourseRepo (db *gorm.DB) repository.CourseRepo{
	return &courseRepo{db}
}



func (cr *courseRepo) FindAll(c  []*model.Course) ([]*model.Course , error)
func (cr *courseRepo) FindPreReqs(ID uint ,c  []*model.Course) ([] *model.Course , error)
func (cr *courseRepo) CreateCourse(c *model.Course) (error){return nil}
func (cr *courseRepo) DeleteCourse(ID uint) (error)