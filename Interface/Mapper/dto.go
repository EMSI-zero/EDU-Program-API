package mapper

import (
	"Basic-Edu-Program/Domain/model"
	"Basic-Edu-Program/UseCase/presenter"
)

type courseDto struct{

}

type coursePresenter struct{}


func NewCoursePresenter() presenter.CoursePresenter{
	return &coursePresenter{}
}

func (cd *coursePresenter) ResponsePreReqs(c []* model.Course) ([]*model.Course , error){
	return nil,nil
}
func (cd *coursePresenter) ResponseOperation(error) (string){
	return ""
}
func (cd *coursePresenter) ResponseCourses(c []*model.Course) ([]*model.Course){
	return nil
}