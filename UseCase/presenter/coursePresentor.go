package presenter

import "Basic-Edu-Program/Domain/model"


type CoursePresenter interface {
	ResponsePreReqs(c []* model.Course) ([]*model.Course , error)
	ResponseOperation(error) (string)
	ResponseCourses(c []*model.Course) ([]*model.Course)
}