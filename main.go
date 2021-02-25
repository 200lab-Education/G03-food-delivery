package main

import (
	"demo/component/appctx"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	sysSecret := os.Getenv("SYSTEM_SECRET")

	dns := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()
	appCtx := appctx.New(db, sysSecret)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	setupRouter(r, appCtx)

	if err := r.Run(); err != nil {
		log.Fatalln(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
