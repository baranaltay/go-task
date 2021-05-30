package controllers

import (
	"net/http"
	"time"

	"github.com/baranaltay/go-task/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var user = models.User{
	ID:       "1",
	Username: "task",
	Password: "task",
}

const SECRET = "my secret"

func Login(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid model")
		return
	}

	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "invalid credentials")
		return
	}

	atClaims := jwt.MapClaims{}
	atClaims["sub"] = user.ID
	atClaims["username"] = user.Username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims).SignedString([]byte(SECRET))

	if err != nil {
		panic("error while generating jwt")
	}

	c.JSON(http.StatusOK, token)
}

func ChangePassword(c *gin.Context) {
	var requestModel models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&requestModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid model")
		return
	}

	if requestModel.OldPassword != user.Password {
		c.JSON(http.StatusBadRequest, "invalid credentials")
		return
	}
	if requestModel.NewPassword != requestModel.NewPasswordConfirmation {
		c.JSON(http.StatusBadRequest, "passwords did not match")
		return
	}

	user.Password = requestModel.NewPassword
	c.Status(http.StatusNoContent)
}
