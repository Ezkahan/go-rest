package main

import (
	"github.com/ezkahan/go-rest/config"
	"github.com/ezkahan/go-rest/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = config.MysqlConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run(":8090")
}
