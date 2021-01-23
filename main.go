package main

import (
	"demo/common"
	"demo/component/appctx"
	"demo/module/note/notemodel"
	"demo/module/note/notetransport/ginnote"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
)

type LoginData struct {
	Username string `json:"username" form:"username"` // tag
	Password string `json:"password" form:"password"`
}

type UpdateNote struct {
	Title   *string `json:"title" form:"title" gorm:"column:title"`
	Content *string `json:"content" form:"content" gorm:"column:content"`
}

func main() {
	dns := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	appCtx := appctx.New(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	notes := v1.Group("/notes")
	{
		// create a new note
		notes.POST("", func(c *gin.Context) {
			var note notemodel.Note

			if err := c.ShouldBind(&note); err != nil {
				c.AbortWithStatusJSON(400, gin.H{
					"message": "bad request",
				})
				return
			}

			if err := db.Create(&note).Error; err != nil {
				log.Println("cant create a note", err)
			}

			c.JSON(http.StatusOK, gin.H{
				"id": note.Id,
			})
			return
		})

		// get a list of notes
		notes.GET("", ginnote.ListNote(appCtx))

		// get a note details by id
		notes.GET("/:note-id", func(c *gin.Context) {
			var note notemodel.Note
			id, _ := strconv.Atoi(c.Param("note-id"))

			if err := db.Table(note.TableName()).
				Where("id = ?", id).
				First(&note).Error; err != nil {
				log.Println("cant read a note", err)
			}

			c.JSON(http.StatusOK, gin.H{"data": note})
			return
		})

		notes.PUT("/:note-id", func(c *gin.Context) {
			var note UpdateNote
			id, _ := strconv.Atoi(c.Param("note-id"))

			if err := c.ShouldBind(&note); err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "some err"})
			}

			if err := db.Table("notes").Where("id = ?", id).Updates(&note).Error; err != nil {
				log.Println("some error when perform update a note", err)
			}
			c.JSON(http.StatusOK, gin.H{"data": true})

		}) // update a note by id
		notes.DELETE("/:note-id", ginnote.DeleteNote(appCtx)) // delete a note by id
	}

	r.GET("/users/:id/notes") // get a list of notes belong to user by id

	r.POST("auth/login") // login API
	r.POST("/cart/checkout")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type fakeDeleteNoteStore struct{}

func (fakeDeleteNoteStore) FindDataWithCondition(condition map[string]interface{}) (*notemodel.Note, error) {
	return &notemodel.Note{
		SQLModel: common.SQLModel{Id: 1, Status: 1},
		Title:    "",
		Content:  "",
	}, nil
}

func (fakeDeleteNoteStore) Delete(id int) error {
	return nil
}
