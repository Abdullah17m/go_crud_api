package response

import (
	"net/http"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, map[string]interface{}{"success": true, "data": data})
}

func BadRequest(c echo.Context, msg string) error {
	return c.JSON(http.StatusBadRequest, map[string]interface{}{"success": false, "error": msg})
}

func InternalError(c echo.Context, msg string) error {
	return c.JSON(http.StatusInternalServerError, map[string]interface{}{"success": false, "error": msg})
}

func ValidationError(c echo.Context, err error) error {
	errors := make(map[string]string)
	for _, e := range err.(validator.ValidationErrors) {
		errors[e.Field()] = e.Tag() // or use e.Translate(translator)
	}
	return c.JSON(http.StatusBadRequest, echo.Map{
		"error":             "Validation failed",
		"validation_errors": errors,
	})
}