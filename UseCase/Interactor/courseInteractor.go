package interactor

import (
	"Basic-Edu-Program/Domain/models/course"
	"Basic-Edu-Program/UseCase/repository"
	"Basic-Edu-Program/UseCase/presenter"
)

type courseInteractor struct{
	CourseRepository repository.CourseRepo 
	CoursePresenter presenter.CoursePresenter
}

type CourseInteractor interface{
	GetCourses (courses []*course.Course) ([]*course.Course, error)
	FindCoursesByTitle (courses []*course.Course , s []string) (  []*course.Course, error)
	FindCoursesByID (  courses []*course.Course , ID []uint) ([]*course.Course , error)
	FindRequirements (courses []*course.Course , ID uint ) ([]*course.Course, error)
	CreateCourse (c []*course.Course) (error)
	DeleteCourse (c []*course.Course) (error)
}

func NewCourseInteractor (cr repository.CourseRepo , cp presenter.CoursePresenter) CourseInteractor {
	return &courseInteractor{cr , cp}
}

