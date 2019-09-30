package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func showSearchPaymentPage(c *gin.Context) {

	c.HTML(http.StatusOK, "search-student.html", gin.H{
		"title":   "Payment Search"})

}

func showPaymentPage(c *gin.Context) {

	c.HTML(http.StatusOK, "payment.html", gin.H{
		"title":   "Payment"})

}