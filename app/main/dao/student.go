package dao

import (
	"container/app/main/models"
	"container/global"
)

type studentDao struct {
}

var StudentDao = &studentDao{}


func (s *studentDao) QueryAll() (students []models.Student) {
	global.DB.Find(&students)
	return
}

func (s *studentDao) QueryByName(name string) (student models.Student) {
	global.DB.Find(&student, "name = ?", name)
	return
}

func (s *studentDao) CreateStudent(student *models.Student) {
	global.DB.Create(student)
}