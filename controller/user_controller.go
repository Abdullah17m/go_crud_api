package controller

import (
	"crud-api/manager"
	"crud-api/request"
	"crud-api/response"
	"strconv"
	"crud-api/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var validate = validator.New()

func CreateUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return response.BadRequest(c, "Invalid request body")
	}

	
	exists, err := manager.IsEmailExists(user.Email)
	if err != nil {
		return response.InternalError(c, "Failed to check email")
	}
	if exists {
		return response.BadRequest(c, "User with this email already exists")
	}

	createdUser, err := manager.CreateUser(request.CreateUserRequest{
		Name:  user.Name,
		Email: user.Email,
	})
	
	if err != nil {
		return response.InternalError(c, "Failed to create user")
	}

	return response.Success(c, createdUser)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	user, err := manager.GetUser(id)
	if err != nil {
		return response.InternalError(c, err.Error())
	}
	return response.Success(c, user)
}

func GetAllUsers(c echo.Context) error {
	// Get query params
	pageQuery := c.QueryParam("page")
	limitQuery := c.QueryParam("limit")

	// Default values
	page := int64(1)
	limit := int64(10)

	if pageQuery != "" {
		if p, err := strconv.ParseInt(pageQuery, 10, 64); err == nil {
			page = p
		}
	}
	if limitQuery != "" {
		if l, err := strconv.ParseInt(limitQuery, 10, 64); err == nil {
			limit = l
		}
	}

	users, err := manager.GetAllUsers(page, limit)
	if err != nil {
		return response.InternalError(c, err.Error())
	}

	return response.Success(c, echo.Map{
		"page":  page,
		"limit": limit,
		"data":  users,
	})
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var req request.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid input")
	}

	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	user, err := manager.UpdateUser(id, req)
	if err != nil {
		return response.InternalError(c, err.Error())
	}
	return response.Success(c, user)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := manager.DeleteUser(id)
	if err != nil {
		return response.InternalError(c, err.Error())
	}
	return response.Success(c, "User deleted")
}
