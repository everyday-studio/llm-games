package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

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
	group.POST("/login", handler.Login)

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

func (h *AuthHandler) Login(c echo.Context) error {
	req := new(domain.LoginRequest)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrResponse(domain.ErrInvalidInput))
	}

	if req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, ErrResponse(domain.ErrInvalidInput))
	}

	ctx := c.Request().Context()
	loginResponse, err := h.authUseCase.Login(ctx, req.Email, req.Password)
	if err == nil {
		cookie := h.createRefreshTokenCookie(
			loginResponse.RefreshToken,
			loginResponse.RefreshTokenExpiration,
		)
		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, loginResponse)
	}

	switch {
	case errors.Is(err, domain.ErrInvalidInput):
		return c.JSON(http.StatusUnauthorized, ErrResponse(errors.New("invalid email or password")))
	default:
		return c.JSON(http.StatusInternalServerError, ErrResponse(domain.ErrInternal))
	}
}

func (h *AuthHandler) createRefreshTokenCookie(tokenValue string, expiration time.Time) *http.Cookie {
	cookieConfig := h.config.Secure.JWT.Cookie

	sameSite := http.SameSiteLaxMode
	switch strings.ToLower(cookieConfig.SameSite) {
	case "strict":
		sameSite = http.SameSiteStrictMode
	case "lax":
		sameSite = http.SameSiteLaxMode
	case "none":
		sameSite = http.SameSiteNoneMode
	}

	cookie := &http.Cookie{
		Name:     refreshTokenCookieName,
		Value:    tokenValue,
		Path:     "/",
		Domain:   cookieConfig.Domain,
		Expires:  expiration,
		Secure:   cookieConfig.Secure,
		HttpOnly: cookieConfig.HTTPOnly,
		SameSite: sameSite,
	}

	return cookie
}
