package interactor

import "Basic-Edu-Program/Domain/models/course"


func (ci * courseInteractor) GetCourses (courses []*course.Course) ([]*course.Course, error)
func (ci * courseInteractor) FindCoursesByTitle (courses []*course.Course , s []string) (  []*course.Course, error)
func (ci * courseInteractor) FindCoursesByID (  courses []*course.Course , ID []uint) ([]*course.Course , error)
func (ci * courseInteractor) FindRequirements (courses []*course.Course , ID uint ) ([]*course.Course, error)
func (ci * courseInteractor) CreateCourse (c []*course.Course) (error)
func (ci * courseInteractor) DeleteCourse (c []*course.Course) (error)