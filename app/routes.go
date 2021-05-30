package app

import (
	"github.com/baranaltay/go-task/controllers"
)

func RegisterRoutes() {
	router.POST("/login", controllers.Login)
	router.POST("/changepassword", controllers.ChangePassword)
}
