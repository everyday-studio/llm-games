package handler

import (
	"errors"
	"net/http"

	"github.com/everyday-studio/ollm/internal/config"
	"github.com/everyday-studio/ollm/internal/domain"
	"github.com/everyday-studio/ollm/internal/middleware"
	"github.com/labstack/echo/v4"
)

const refreshTokenCookieName = "refresh_token"

type AuthHandler struct {
	authUseCase domain.AuthUsecase
	config      *config.Config
}

func NewAuthHandler(e *echo.Echo, authUseCase domain.AuthUsecase, config *config.Config) *AuthHandler {
	handler := &AuthHandler{
		authUseCase: authUseCase,
		config:      config,
	}

	group := e.Group("/auth", middleware.AllowRoles(domain.RolePublic))
	group.POST("/signup", handler.SignUpUser)

	return handler
}

func (h *AuthHandler) SignUpUser(c echo.Context) error {
	req := new(domain.SignUpRequest)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrResponse(domain.ErrInvalidInput))
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrResponse(domain.ErrInvalidInput))
	}

	user := &domain.User{
		Name:     "TESTUSER", //TODOs
		Email:    req.Email,
		Password: req.Password,
		Role:     domain.RoleUser,
	}

	ctx := c.Request().Context()
	createdUser, err := h.authUseCase.SignUpUser(ctx, user)
	if err == nil {
		return c.JSON(http.StatusCreated, domain.SignUpResponse{
			ID:    createdUser.ID,
			Name:  createdUser.Name,
			Email: createdUser.Email,
		})
	}

	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		return c.JSON(http.StatusBadRequest, ErrResponse(err))
	case errors.Is(err, domain.ErrAlreadyExists):
		return c.JSON(http.StatusConflict, ErrResponse(err))
	default:
		return c.JSON(http.StatusInternalServerError, ErrResponse(domain.ErrInternal))
	}
}
