package service

import (
	"errors"

	"github.com/EricRider/student-manager/internal/model"
)

type StudentService struct {
	students []model.Student
}

// AddStudent 添加学生
func (s *StudentService) AddStudent(student model.Student) error {
	s.students = append(s.students, student)
	return nil
}

// ListStudents 返回所有学生
func (s *StudentService) ListStudents() []model.Student {
	return s.students
}
func (s *StudentService) DeleteStudent(id int) error {

	for i, student := range s.students {

		if student.ID == id {

			s.students = append(s.students[:i], s.students[i+1:]...)

			return nil
		}
	}

	return errors.New("student not found")
}
