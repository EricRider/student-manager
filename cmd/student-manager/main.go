package main

import (
	"github.com/EricRider/student-manager/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	h := handler.NewStudentHandler()

	r.GET("/students", h.ListStudents)

	r.POST("/students", h.AddStudent)

	r.Run(":8080")
}
