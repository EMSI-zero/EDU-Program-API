package interactor

import (
	"Basic-Edu-Program/Domain/models/course"
	"Basic-Edu-Program/UseCase/repository"
	"Basic-Edu-Program/UseCase/presenter"
)

type userInteractor struct{
	CourseRepository repository.CourseRepo 
	CoursePresenter presenter.CoursePresenter
}

type UserInteractor interface{
	GetCourses () ([]*course.Course, error)
	FindCourses (c  []*course.Course) (  []*course.Course, error)
	GetCredit (c *course.Course) (uint, error)
	FindRequirements (c *course.Course) ([]*course.Course, error)
	CreateCourse (c *course.Course) (error)
	DeleteCourse (c *course.Course) (error)
}