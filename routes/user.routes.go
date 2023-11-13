package routes

import (
	user_controller "kosei-lms-api/controllers/user_controller"
	"kosei-lms-api/middleware"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController user_controller.UserController
}

func NewRouteUserController(userController user_controller.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.userController.GetMe)
	router.PUT("/update", middleware.DeserializeUser(), uc.userController.UpdateUserPhoto)
}
