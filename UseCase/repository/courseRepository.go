package repository

import "Basic-Edu-Program/Domain/model"

type CourseRepo interface{
	FindAll(c  []*model.Course) ([]*model.Course , error)
	FindPreReqs(ID uint ,c  []*model.Course) ([] *model.Course , error)
	CreateCourse(c *model.Course) (error)
	DeleteCourse(ID uint) (error)
}