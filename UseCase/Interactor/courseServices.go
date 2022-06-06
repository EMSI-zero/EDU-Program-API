package interactor

import "Basic-Edu-Program/Domain/models/course"


func (ci * courseInteractor) GetCourses (courses []*course.Course) ([]*course.Course, error){
	cs, err := ci.CourseRepository.findAll(courses)
	if err != nil {
		return nil, err
	}

	return ci.CoursePresenter.ResponseCourses(cs), nil
}


func (ci * courseInteractor) FindCoursesByTitle (courses []*course.Course , s []string) (  []*course.Course, error)
func (ci * courseInteractor) FindCoursesByID (  courses []*course.Course , ID []uint) ([]*course.Course , error)

func (ci * courseInteractor) FindRequirements (courses []*course.Course , ID uint ) ([]*course.Course, error){
	cs , err := ci.CourseRepository.findPreReqs(courses)

	if err != nil {
		return nil, err
	}

	return ci.CoursePresenter.ResponsePreReqs(cs)
}


func (ci * courseInteractor) CreateCourse (c []*course.Course) (error)
func (ci * courseInteractor) DeleteCourse (c []*course.Course) (error)