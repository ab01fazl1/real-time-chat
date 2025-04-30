package user

import (
	"log"
	"net/http"

	auth "github.com/ab01fazl1/real-time-chat/internal/auth"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// TODO: change to user req struct
	// get req body
	var loginReq User
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hit the db to check if user exists
	hashpass, err := get_one_user(c, loginReq.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user doesn't exists, error": err.Error()})
		return
	}

	// compare hash
	err = CheckPassHash(hashpass, loginReq.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"incorrect password, error": err.Error()})
		return
	}

	// generate jwt access token
	accessToken, err := auth.CreateAccessToken(loginReq.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"failed to create access token, error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, accessToken)

}

func Register(c *gin.Context) {
	// get req body
	var signupRequest *SignupRequest
	if err := c.ShouldBind(&signupRequest); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	//TODO: validation function should be implemented later


	// hit the db to check if user exists
	exists, err := userExists(c, signupRequest.Username)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if exists {
		c.String(http.StatusBadRequest, "user already exists")
		return
	}
	
	// hash the pass
	hashedpass, err := HashPass(signupRequest.Password)
	if err != nil {
		log.Println(err.Error())
	}

	// create user
	err = createUser(c, signupRequest.Username, hashedpass)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	
	c.String(http.StatusAccepted, "user with username: %v and password %v created", signupRequest.Username, signupRequest.Password)
	
}