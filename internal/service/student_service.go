package service

import (
	"errors"
	"fmt"

	"github.com/EricRider/student-manager/internal/model"
	"github.com/EricRider/student-manager/internal/storage"
)

type StudentService struct {
	students []model.Student
	storage  *storage.JSONStorage
}

func (s *StudentService) AddStudent(student model.Student) error {
	s.students = append(s.students, student)

	fmt.Println("Saving students:", s.students)

	return s.storage.Save(s.students)
}

// ListStudents 返回所有学生
func (s *StudentService) ListStudents() []model.Student {
	return s.students
}
func (s *StudentService) DeleteStudent(id int) error {

	for i, student := range s.students {

		if student.ID == id {

			s.students = append(s.students[:i], s.students[i+1:]...)

			return s.storage.Save(s.students)
		}
	}

	return errors.New("student not found")
}
func NewStudentService() *StudentService {

	s := &StudentService{
		storage: storage.NewJSONStorage("students.json"),
	}

	students, err := s.storage.Load()

	if err != nil {
		panic(err)
	}

	s.students = students

	return s
}
