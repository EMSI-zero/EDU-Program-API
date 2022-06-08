package repository

import "Basic-Edu-Program/Domain/models/course"


func findAll(c  []*course.Course) ([]*course.Course , error)
func findPreReqs(ID uint ,c  *[]course.Course) ([] *course.Course , error)
func createCourse(c *course.Course) (error)
func deleteCourse(ID uint) (error)