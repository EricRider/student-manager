package handler

import (
	"net/http"

	"github.com/EricRider/student-manager/internal/model"
	"github.com/EricRider/student-manager/internal/service"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	service *service.StudentService
}

func NewStudentHandler() *StudentHandler {
	return &StudentHandler{
		service: service.NewStudentService(),
	}
}

// GET /students
func (h *StudentHandler) ListStudents(c *gin.Context) {

	students := h.service.ListStudents()

	c.JSON(http.StatusOK, students)
}

// POST /students
func (h *StudentHandler) AddStudent(c *gin.Context) {

	var student model.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	err := h.service.AddStudent(student)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student added successfully",
	})
}
