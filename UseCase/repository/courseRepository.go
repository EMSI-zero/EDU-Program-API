package repository

import "Basic-Edu-Program/Domain/models/course"

type CourseRepo interface{
	findAll(c  []*course.Course) ([]*course.Course , error)
	findPreReqs(ID uint ,c  *[]course.Course) ([] *course.Course , error)
	createCourse(c *course.Course) (error)
	deleteCourse(ID uint) (error)
}