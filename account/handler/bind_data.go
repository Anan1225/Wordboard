package handler

import (
	"log"

	"github.com/Anan1225/wordboard/account/model/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// used to helo extract validation errors
type invalidArgument struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Tag   string `json:"tag"`
	Param string `json:"param"`
}

// bindData is helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		log.Printf("Error binding data: %+v\n", err)

		if errs, ok := err.(validator.ValidationErrors); ok {
			// could probably extract this, it is also in middleware_auth_user
			var invalidArgs []invalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, invalidArgument{
					err.Field(),
					err.Value().(string),
					err.Tag(),
					err.Param(),
				})
			}

			err := apperrors.NewBadRequest("Invaild request parameters. See invalidArgs")

			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}

		// later we'll add code for validating max body size here!

		// if we aren't able to properly extract validation errors,
		// we will fallback and return an internal server error
		fallback := apperrors.NewInternal()

		c.JSON(fallback.Status(), gin.H{"error": fallback})
		return false
	}
	return true
}