package services

import (
	"container/app/main/dao"
	"container/app/main/models"
	"container/app/main/models/request"
	"errors"
)

type studentService struct {}

var StudentService = &studentService{}


func (s *studentService) CreateStudent(studentRequest *request.Student) error {
	student := models.Student{
		Name: studentRequest.Name,
		StudentId: studentRequest.StudentId,
		Grade: studentRequest.Grade,
	}

	dao.StudentDao.CreateStudent(&student)
	return nil
}


func (s *studentService) FindStudentByName(name string) (student models.Student, err error) {
	student = dao.StudentDao.QueryByName(name)
	if student.ID == 0 {
		return student, errors.New("该学生不存在")
	}
	return
}

func (s *studentService) FindAll() (students []models.Student) {
	return dao.StudentDao.QueryAll()
}