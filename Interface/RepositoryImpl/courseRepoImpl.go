package repository

import (
	"Basic-Edu-Program/Domain/models/course"
	"Basic-Edu-Program/UseCase/repository"

	"gorm.io/gorm"
)


type courseRepo struct {
	db *gorm.DB
}

func NewCourseRepo (db *gorm.DB) repository.CourseRepo{
	return &courseRepo{db}
}

func (cr *courseRepo) FindAll(c  []*course.Course) ([]*course.Course , error)
func (cr *courseRepo) FindPreReqs(ID uint ,c  []*course.Course) ([] *course.Course , error)
func (cr *courseRepo) CreateCourse(c *course.Course) (error){return nil}
func (cr *courseRepo) DeleteCourse(ID uint) (error)