package presenter

import "Basic-Edu-Program/Domain/models/course"


type CoursePresenter interface {
	ResponsePreReqs(c []* course.Course) ([]*course.Course , error)
	ResponseCoursePath(c [][]*course.Course) ([][]*course.Course, error)
	ResponseOperation(error) (string)
	ResponseCourses(c []*course.Course) ([]*course.Course)
}