package ui

import (
	"fmt"

	"github.com/EricRider/student-manager/internal/model"
	"github.com/EricRider/student-manager/internal/service"
)

// Start 启动菜单
func Start() {
	studentService := service.NewStudentService()
	for {

		fmt.Println("======================")
		fmt.Println(" Student Manager ")
		fmt.Println("======================")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Student")
		fmt.Println("3. Delete Student")
		fmt.Println("0. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:

			var student model.Student

			fmt.Print("Input ID: ")
			fmt.Scanln(&student.ID)

			fmt.Print("Input Name: ")
			fmt.Scanln(&student.Name)

			fmt.Print("Input Age: ")
			fmt.Scanln(&student.Age)

			fmt.Print("Input Email: ")
			fmt.Scanln(&student.Email)

			err := studentService.AddStudent(student)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Student Added Successfully!")
			}

		case 2:
			fmt.Println("==========================================================")
			fmt.Printf("%-8s %-15s %-8s %-25s\n", "ID", "Name", "Age", "Email")
			fmt.Println("==========================================================")

			for _, student := range studentService.ListStudents() {
				fmt.Printf("%-8d %-15s %-8d %-25s\n",
					student.ID,
					student.Name,
					student.Age,
					student.Email,
				)

			}

			fmt.Println("==========================================================")

		case 3:
			var id int
			fmt.Println("Input ID:")
			fmt.Scanln(&id)
			err := studentService.DeleteStudent(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Student Deleted Successfully!")
			}

		case 0:
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Invalid Choice")
		}

		fmt.Println()
	}

}
