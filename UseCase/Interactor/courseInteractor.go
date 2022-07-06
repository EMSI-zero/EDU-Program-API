package interactor

import (
	"Basic-Edu-Program/Domain/model"
	"Basic-Edu-Program/UseCase/repository"
	"Basic-Edu-Program/UseCase/presenter"
)

type courseInteractor struct{
	CourseRepository repository.CourseRepo 
	CoursePresenter presenter.CoursePresenter
}

type CourseInteractor interface{
	GetCourses (courses []*model.Course) ([]*model.Course, error)
	FindCoursesByTitle (courses []*model.Course , s []string) (  []*model.Course, error)
	FindCoursesByID (  courses []*model.Course , ID []uint) ([]*model.Course , error)
	FindRequirements (courses []*model.Course , ID uint ) ([]*model.Course, error)
	CreateCourse (c []*model.Course) (error)
	DeleteCourse (c []uint) (error)
}

func NewCourseInteractor (cr repository.CourseRepo , cp presenter.CoursePresenter) CourseInteractor {
	return &courseInteractor{cr , cp}
}

