package handler

import (
	"baby_starter/app"
	"baby_starter/database/model"
	"baby_starter/util"
	"errors"
	"net/http"

	z "github.com/Oudwins/zog"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PostUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func PostUser(c echo.Context) error {
	var req PostUserRequest
	if err, ok := ParseReq(c, z.Struct(z.Shape{
		"email": z.String().
			Trim().
			Email(z.Message("Please provide a valid email address")).
			Max(255, z.Message("Email must be at most 255 characters long")).
			Required(z.Message("Email is required")),
		"password": z.String().
			Trim().
			Min(8, z.Message("Password must be at least 8 characters long")).
			Max(255, z.Message("Password must be at most 255 characters long")).
			Required(z.Message("Password is required")),
		"name": z.String().
			Trim().
			Min(2, z.Message("Name must be at least 2 characters long")).
			Max(255, z.Message("Name must be at most 255 characters long")).
			Required(z.Message("Name is required")),
	}), &req); !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err,
		})
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		app.LOG.Error().Err(err).Msg("failed to hash password")
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to hash password",
		})
	}

	// check if user exists
	if _, err := gorm.G[model.User](app.DB).Where("email = ?", req.Email).First(c.Request().Context()); err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "User already exists",
		})
	}

	user := model.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
		IsAdmin:  false,
	}

	if err := gorm.G[model.User](app.DB).Create(c.Request().Context(), &user); err != nil {
		app.LOG.Error().Err(err).Msg("failed to create user")
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, user)
}
