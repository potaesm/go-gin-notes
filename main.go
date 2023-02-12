package main

import (
	"log"
	"net/http"

	"github.com/potaesm/go-gin-notes/controller"
	controllerHelper "github.com/potaesm/go-gin-notes/controller/helper"
	"github.com/potaesm/go-gin-notes/middleware"
	"github.com/potaesm/go-gin-notes/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.Static("/vendor", "./static/vendor")

	r.LoadHTMLGlob("template/**/*")

	model.ConnectDatabase()
	model.DBMigrate()

	// Sessions Init
	store := memstore.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("notes", store))

	r.Use(middleware.AuthenticateUser())

	// Route Group - Notes
	notes := r.Group("/notes")
	{
		notes.GET("/", controller.NotesIndex)
		notes.GET("/new", controller.NotesNew)
		notes.POST("/", controller.NotesCreate)
		notes.GET("/:id", controller.NotesShow)
		notes.GET("/edit/:id", controller.NotesEditPage)
		notes.POST("/:id", controller.NotesUpdate)
		notes.DELETE("/:id", controller.NotesDelete)
	}

	r.GET("/login", controller.LoginPage)
	r.GET("/signup", controller.SignupPage)

	r.POST("/signup", controller.Signup)
	r.POST("/login", controller.Login)
	r.POST("/logout", controller.Logout)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home/index.html", controllerHelper.SetPayload(c, gin.H{
			"title":     "Notes application",
			"logged_in": controllerHelper.IsUserLoggedIn(c),
		}))
	})

	log.Println("Server started!")
	r.Run() // Default Port 8080
}
