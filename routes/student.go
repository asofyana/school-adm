package routes

import (
	"net/http"

	"github.com/asofyana/school-adm/models"
	"github.com/gin-gonic/gin"
)

func searchStudentApi(c *gin.Context) {

	var fullName = c.PostForm("fullName")
	var grade = c.PostForm("grade")
	var studentList = models.SearchStudent(fullName, grade)

	c.JSON(http.StatusOK, studentList)

}
