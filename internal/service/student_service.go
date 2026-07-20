package service

import "github.com/EricRider/student-manager/internal/model"

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
