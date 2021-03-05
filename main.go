package main

import (
	"demo/component/appctx"
	"demo/component/uploadprovider"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	sysSecret := os.Getenv("SYSTEM_SECRET")
	dns := os.Getenv("DB_CONN")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3Secret := os.Getenv("S3_SECRET")
	s3Domain := fmt.Sprintf("https://%s", os.Getenv("S3_DOMAIN"))

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	upProvider := uploadprovider.NewS3Provider(s3BucketName, s3Region,
		s3APIKey, s3Secret, s3Domain)

	appCtx := appctx.New(db, sysSecret, upProvider)

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
