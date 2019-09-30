package routes

import (
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {

	router.GET("/", showLoginPage)

	router.POST("/login", performLogin)

	router.GET("/home", showHomePage)

	router.GET("/payment-search", showSearchPaymentPage)

	router.GET("/payment", showPaymentPage)

	router.POST("/search-student-api", searchStudentApi)

}
