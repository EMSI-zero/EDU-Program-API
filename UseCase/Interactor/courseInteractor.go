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
	GetCourses (courses []*model.Course) ([]*presenter.DTO, error)
	FindCoursesByTitle (courses []*model.Course , s []string) (  []*presenter.DTO, error)
	FindCoursesByID (  courses []*model.Course , ID []uint) ([]*presenter.DTO , error)
	FindRequirements (courses []*model.Course , ID uint ) ([]*presenter.DTO, error)
	CreateCourse (c []*model.Course) (error)
	DeleteCourse (c []uint) (error)
}

func NewCourseInteractor (cr repository.CourseRepo , cp presenter.CoursePresenter) CourseInteractor {
	return &courseInteractor{cr , cp}
}

