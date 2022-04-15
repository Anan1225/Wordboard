package handler

import (
	"net/http"
	"time"

	"github.com/Anan1225/wordboard/account/handler/middleware"
	"github.com/Anan1225/wordboard/account/model"
	"github.com/Anan1225/wordboard/account/model/apperrors"
	"github.com/gin-gonic/gin"
)

// Handler struct holds required services for handler to function
type Handler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// Config will hold services that will eventually be injected into this
// handler layer on handler initialization
type Config struct {
	R               *gin.Engine
	UserService     model.UserService
	TokenService    model.TokenService
	BaseURL         string
	TimeoutDuration time.Duration
}

// NewHandler initializes the handler with required injected services along with http routes
// Does not return as it deals directly with a reference to the gin Engine
func NewHandler(c *Config) {
	// Create a handler (which will later have injected services)
	h := &Handler{
		UserService:  c.UserService,
		TokenService: c.TokenService,
	}
	// Create a group, or base url for all routes
	// g := c.R.Group("/api/account")
	g := c.R.Group(c.BaseURL)

	// Wish I had thought this through better!
	if gin.Mode() != gin.TestMode {
		g.Use(middleware.Timeout(c.TimeoutDuration, apperrors.NewServiceUnavailable()))
		g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
		g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
	} else {
		g.GET("/me", h.Me)
		g.POST("/signout", h.Signout)
	}

	// g.GET("/me", h.Me)
	g.POST("/signup", h.Signup)
	g.POST("/signin", h.Signin)
	g.POST("/tokens", h.Tokens)
	g.POST("/image", h.Image)
	g.DELETE("/image", h.DeleteImage)
	g.PUT("/details", h.Details)

	// g.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"hello": "space persons",
	// 	})
	// })
}

// Me handler calls services for getting
// a user's details
// func (h *Handler) Me(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"hello": "it's me",
// 	})
// }

// Tokens handler
func (h *Handler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's token",
	})
}

// Image handler
func (h *Handler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *Handler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteimage",
	})
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
