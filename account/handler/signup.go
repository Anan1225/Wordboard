package handler

import (
	"github.com/gin-gonic/gin"
)

// signupReq is not exported
type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *Handler) Signup(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"hello": "it's signup",
	// })

	// define a variable to which we'll bind incoming
	// json body, {email, passeord}
	var req signupReq

}
