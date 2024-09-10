package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Megidy/To-Do-List-Api/pkj/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	var user models.User
	//New request with new user
	err := c.ShouldBindJSON(&user)

	//checking if user already exist
	value, response := models.IsSignedUp(user)
	if value {
		c.JSON(http.StatusFound, gin.H{
			"error": response,
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didn't read body",
		})
	}
	//hashin password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "filed to hash password",
		})
		return
	}

	NewUser := models.User{
		NickName: user.NickName,
		Email:    user.Email,
		Password: string(hash),
	}

	//adding new user to database
	_, err = models.CreateUser(&NewUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to craete new user",
		})
		return
	}
	//response
	c.JSON(http.StatusOK, gin.H{})

}

func LogIn(c *gin.Context) {
	var NewUserRequest models.User

	//json request
	err := c.ShouldBindJSON(&NewUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didn't read body",
		})
		return
	}

	//finding in database user
	user, err := models.FindUserByEmail(NewUserRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password ",
		})
		return
	}
	//comparing 2 passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.Password),
		[]byte(NewUserRequest.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password ",
		})
		return
	}
	//jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	//write jwt token in cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*25*30, "", "", false, true)

	//response
	c.JSON(http.StatusOK, gin.H{
		"Message": "you logged in ",
	})

}
