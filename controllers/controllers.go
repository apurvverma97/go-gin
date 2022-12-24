package controllers

import (
	"master-go/crud-apis/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func new(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context)  {
	ctx.JSON(200, "")
}

func (uc *UserController) GetUser(ctx *gin.Context)  {
	ctx.JSON(200, "")
}

func (uc *UserController) GetAll(ctx *gin.Context)  {
	ctx.JSON(200, "")
}

func (uc *UserController) UpdateUser(ctx *gin.Context)  {
	ctx.JSON(200, "")
}

func (uc *UserController) DeleteUser(ctx *gin.Context)  {
	ctx.JSON(200, "")
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup){
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/get/:name", uc.GetUser)
	userRoute.GET("/getall", uc.GetAll)
	userRoute.PATCH("/update", uc.UpdateUser)
	userRoute.DELETE("/delete", uc.DeleteUser)
}