package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/potaesm/go-gin-notes/model"
)

func GetUserFromRequest(c *gin.Context) *model.User {
	// Get user
	userID := c.GetUint64("user_id")
	var currentUser *model.User
	if userID > 0 {
		currentUser = model.UserFind(userID)
	} else {
		currentUser = nil
	}
	return currentUser
}

func IsUserLoggedIn(c *gin.Context) bool {
	return (c.GetUint64("user_id") > 0)
}

func SetPayload(c *gin.Context, h gin.H) gin.H {
	email := c.GetString("email")
	if len(email) > 0 {
		h["email"] = email
	}
	return h
}
