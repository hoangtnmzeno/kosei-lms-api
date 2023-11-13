package routes

import (
	auth_controller "kosei-lms-api/controllers/auth_controller"
	"kosei-lms-api/middleware"

	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController auth_controller.AuthController
}

func NewAuthRouteController(authController auth_controller.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
