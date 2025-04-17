package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var json User
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, json)
}

func Register(c *gin.Context) {
	// get req body
	var signupRequest *SignupRequest
	if err := c.ShouldBind(&signupRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	// validation function should be implemented later


	// hit the db to check if user exists
	exists, err := userExists(c, signupRequest.Username)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	if exists {
		c.String(http.StatusBadRequest, "user already exists, err: %v", err.Error())
	}

	// create user
	err = createUser(c, signupRequest.Username, signupRequest.Password)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	
	c.String(http.StatusAccepted, "user with username: %v and password %v created", signupRequest.Username, signupRequest.Password)
	
}