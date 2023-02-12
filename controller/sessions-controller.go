package controller

import (
	"fmt"
	"net/http"

	"github.com/potaesm/go-gin-notes/helper"
	"github.com/potaesm/go-gin-notes/model"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{},
	)
}

func SignupPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"home/signup.html",
		gin.H{},
	)
}

func Signup(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirmPassword := c.PostForm("confirm_password")

	// Check if email already exists
	available := model.UserCheckAvailability(email)
	fmt.Println(available)
	if !available {
		c.HTML(
			http.StatusIMUsed,
			"home/signup.html",
			gin.H{
				"alert": "Email already exists!",
			},
		)
		return
	}
	if password != confirmPassword {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Password missmatch!",
			},
		)
		return
	}
	user := model.UserCreate(email, password)
	if user.ID == 0 {
		c.HTML(
			http.StatusNotAcceptable,
			"home/signup.html",
			gin.H{
				"alert": "Unable to create user!",
			},
		)
	} else {
		// Signup successful, set session
		helper.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	}
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	user := model.UserCheck(email, password)
	if user != nil {
		// Set session
		helper.SessionSet(c, user.ID)
		c.Redirect(http.StatusMovedPermanently, "/")
	} else {
		c.HTML(
			http.StatusOK,
			"home/login.html",
			gin.H{
				"alert": "Email and/or password mismatch!",
			},
		)
	}
}

func Logout(c *gin.Context) {
	// Clear the session
	helper.SessionClear(c)
	c.HTML(
		http.StatusOK,
		"home/login.html",
		gin.H{
			"alert": "Logged out",
		},
	)
}
