package handler

import (
	"log"
	"net/http"

	"github.com/Anan1225/wordboard/account/model"
	"github.com/Anan1225/wordboard/account/model/apperrors"
	"github.com/gin-gonic/gin"
)

// Me handler calls services for getting
// a user's details

func (h *Handler) Me(c *gin.Context) {
	// A *model.User will eventually be added to context in middleware
	user, exists := c.Get("user")

	// This shouldn't happen, as our middleware ought to throw an error.
	// We can also use "MustGet" to get the key or panic
	if !exists {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
		err := apperrors.NewInternal()
		c.JSON(err.Status(), gin.H{
			"error": err,
		})

		return
	}

	uid := user.(*model.User).UID

	// use the Request Context
	ctx := c.Request.Context()

	// gin.Context satisfies go's context.Context interface
	u, err := h.UserService.Get(ctx, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := apperrors.NewNotFound("user", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})
}
