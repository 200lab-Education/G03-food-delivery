package main

import (
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

type Note struct {
	Id      int    `json:"id,omitempty" gorm:"column:id"`
	Title   string `json:"title" form:"title" gorm:"column:title"`
	Content string `json:"content" form:"content" gorm:"column:content"`
}

type UpdateNote struct {
	Title   *string `json:"title" form:"title" gorm:"column:title"`
	Content *string `json:"content" form:"content" gorm:"column:content"`
}

func (Note) TableName() string {
	return "notes"
}

func main() {

	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// "deliveryfood:deliveryfood@tcp(127.0.0.1:3303)/deliveryfood?charset=utf8mb4&parseTime=True&loc=Local"
	dns := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	log.Println("connected to db.", db)

	//var note Note
	//note.Id = 1
	//note.Title = "hoc lap trinh golang"
	//note.Content = "thuc hanh bai 4"
	//
	//if err := db.Create(&note).Error; err != nil {
	//	log.Println("cant create a note", err)
	//}

	//log.Println("before read from db", note)
	//
	//if err := db.Table(note.TableName()).
	//	Where("id = ?", 1).
	//	First(&note).Error; err != nil {
	//	log.Println("cant read a note", err)
	//}
	//
	//log.Println("after read from db", note.Id, note.Title, note.Content)

	//if err := db.Table(note.TableName()).
	//	Where("id = ?", 1).
	//	Update("content", "hoc lap trinh golang vao toi thu 5").Error; err != nil {
	//	log.Println("cant update a note", err)
	//}

	//if err := db.Table(note.TableName()).
	//	Where("id = ?", 1).
	//	Delete(nil).Error; err != nil {
	//	log.Println("cant delete a note", err)
	//}
	log.Println("deleted a note")

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
		// create a new note
		notes.POST("", func(c *gin.Context) {
			var note Note

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
		notes.GET("", func(c *gin.Context) {
			var notes []Note

			if err := db.Table("notes").
				Offset(1).
				Limit(1).
				Find(&notes).Error; err != nil {
				log.Println("cant read a note", err)
			}

			c.JSON(http.StatusOK, gin.H{"data": notes})
			return
		})

		// get a note details by id
		notes.GET("/:note-id", func(c *gin.Context) {
			var note Note
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
		notes.DELETE("/:id") // delete a note by id
	}

	r.GET("/users/:id/notes") // get a list of notes belong to user by id

	r.POST("auth/login") // login API
	r.POST("/cart/checkout")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
