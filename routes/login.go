package routes

import (
	"net/http"

	"github.com/asofyana/school-adm/middleware"
	"github.com/gin-gonic/gin"
)

func showLoginPage(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", nil)

}

func performLogin(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")

	authStatus, authMessage := middleware.AuthUser(username, password)

	if authStatus {
		c.HTML(http.StatusOK, "home.html", nil)
	} else {
		c.HTML(http.StatusOK, "login.html", gin.H{"message": authMessage})
	}

}

func logout(c *gin.Context) {

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
