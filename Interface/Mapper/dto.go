package mapper

import (
	"Basic-Edu-Program/Domain/model"
	"Basic-Edu-Program/UseCase/presenter"
)

type courseDto struct{
	ID 				uint		`json: "id"`
	Title			uint		`json: "title"`
	Credit			uint		`json: "credit"`
	Description		string		`json: "decription"`
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