package repository

import "Basic-Edu-Program/Domain/models/course"

type CourseRepo interface{
	findAll(c  []*course.Course) ([]*course.Course , error)
	getCredit(c * course.Course) ( int , error)
	findPreReqs(c  *course.Course) ([][] *course.Course , error)
	createCourse(c *course.Course) (error)
	deleteCourse(c *course.Course) (error)
}