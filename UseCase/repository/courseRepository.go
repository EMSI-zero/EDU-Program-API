package repository

import "Basic-Edu-Program/Domain/models/course"

type CourseRepo interface{
	FindAll(c  []*course.Course) ([]*course.Course , error)
	FindPreReqs(ID uint ,c  []*course.Course) ([] *course.Course , error)
	CreateCourse(c *course.Course) (error)
	DeleteCourse(ID uint) (error)
}