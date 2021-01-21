package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoginData struct {
	Username string `json:"username" form:"username"` // tag
	Password string `json:"password" form:"password"`
}

type Note struct {
	Id        int    `json:"id,omitempty"`
	Title     string `json:"title"`
	LoginData `json:"a"`
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/demo", func(c *gin.Context) {
		c.JSON(http.StatusOK, Note{})
	})

	r.POST("/demo", func(c *gin.Context) {
		var data LoginData

		c.ShouldBind(&data)

		c.JSON(http.StatusOK, data)
	})

	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	{
		notes.POST("") // create a new note
		notes.GET("")  // get a list of notes
		notes.GET("/:note-id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("note-id"))

			c.JSON(http.StatusOK, gin.H{"id": id})
		}) // get a note details by id
		notes.PUT("/:id")    // update a note by id
		notes.DELETE("/:id") // delete a note by id
	}

	r.GET("/users/:id/notes") // get a list of notes belong to user by id

	r.POST("auth/login") // login API
	r.POST("/cart/checkout")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
