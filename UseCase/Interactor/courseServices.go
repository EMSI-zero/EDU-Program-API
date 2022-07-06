package interactor

import "Basic-Edu-Program/Domain/model"


func (ci * courseInteractor) GetCourses (courses []*model.Course) ([]*model.Course, error){
	cs, err := ci.CourseRepository.FindAll(courses)
	if err != nil {
		return nil, err
	}

	return ci.CoursePresenter.ResponseCourses(cs), nil
}


func (ci * courseInteractor) FindCoursesByTitle (courses []*model.Course , s []string) (  []*model.Course, error)
func (ci * courseInteractor) FindCoursesByID (  courses []*model.Course , ID []uint) ([]*model.Course , error)

func (ci * courseInteractor) FindRequirements (courses []*model.Course , ID uint ) ([]*model.Course, error){
	cs , err := ci.CourseRepository.FindPreReqs(ID, courses)

	if err != nil {
		return nil, err
	}

	return ci.CoursePresenter.ResponsePreReqs(cs)
}


func (ci * courseInteractor) CreateCourse (c []*model.Course) (error){

	for _ ,course := range c {
		err := ci.CourseRepository.CreateCourse(course)
		if err != nil {
			return err
		}
	}

	return nil
}


func (ci * courseInteractor) DeleteCourse (c []uint ) (error){
	
	for _ , ID := range c {

		err := ci.CourseRepository.DeleteCourse(ID)
		if err != nil {
			return err
		}
	}
	return nil
}