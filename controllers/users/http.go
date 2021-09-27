package users

import (
	"net/http"
	"strconv"
	"user-crud-ca/business/users"
	"user-crud-ca/controllers/users/request"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService users.Service
}

func NewUserController(userCtrl users.Service) *UserController{
	return &UserController{
		userService: userCtrl,
	}
}

func (ctrl *UserController) GetByUserID(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	
	res, err := ctrl.userService.GetByUserID(ctx, convId)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (ctrl *UserController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	res, err := ctrl.userService.GetAll(ctx)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (ctrl *UserController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.User
	c.Bind(&req)

	res, err := ctrl.userService.Store(ctx, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (ctrl *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.User
	c.Bind(&req)
	id := c.Param("id")
	convId, err := strconv.Atoi(id)

	res, err := ctrl.userService.Update(ctx, convId, req.ToDomain())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}

	return c.JSON(http.StatusOK, res)
}

func (ctrl *UserController) Delete(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param("id")
	convId, err := strconv.Atoi(id)
	

	_, err = ctrl.userService.GetByUserID(ctx, convId)
	
	if err != nil{
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	} 

	err = ctrl.userService.Delete(ctx, convId)
	if err != nil{
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
		})
	}
	
	return c.JSON(http.StatusOK,"deleted")
}