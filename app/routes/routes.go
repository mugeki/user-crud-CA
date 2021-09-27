package routes

import (
	"user-crud-ca/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (ctrlList *ControllerList) RouteRegister(e *echo.Echo){
	e.GET("/user/:id",ctrlList.UserController.GetByUserID)
	e.PUT("/user/:id",ctrlList.UserController.Update)
	e.DELETE("/user/:id",ctrlList.UserController.Delete)
	e.GET("/user",ctrlList.UserController.GetAll)
	e.POST("/user",ctrlList.UserController.Store)
}