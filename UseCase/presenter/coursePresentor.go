package presenter

import "Basic-Edu-Program/Domain/model"




type DTO interface{}


type CoursePresenter interface {
	ResponsePreReqs(c []* model.Course) ([]*DTO , error)
	ResponseOperation(error) (*DTO)
	ResponseCourses(c []*model.Course) ([]*DTO)
}