package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kavindu-prabhashitha/go-crud-gin-mongodb/controllers"
	"github.com/kavindu-prabhashitha/go-crud-gin-mongodb/middleware"
	"github.com/kavindu-prabhashitha/go-crud-gin-mongodb/services"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middleware.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
