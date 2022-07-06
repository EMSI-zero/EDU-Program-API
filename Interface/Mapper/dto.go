package mapper

import (
	"Basic-Edu-Program/Domain/model"
	"Basic-Edu-Program/UseCase/presenter"
)

type courseDto struct{
	ID 				uint		`json: "id,omitempty"`
	Title			uint		`json: "title,omitempty"`
	Credit			uint		`json: "credit,omitempty"`
	Description		string		`json: "decription,omitempty"`
}

type coursePresenter struct{}


func NewCoursePresenter() presenter.CoursePresenter{
	return &coursePresenter{}
}

func (cd *coursePresenter) ResponsePreReqs(c []* model.Course) ([]*presenter.DTO , error){
	return nil,nil
}
func (cd *coursePresenter) ResponseOperation(error) (*presenter.DTO){
	return nil
}
func (cd *coursePresenter) ResponseCourses(c []*model.Course) ([]*presenter.DTO){
	return nil
}