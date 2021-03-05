package main

import (
	"demo/component/appctx"
	"demo/middleware"
	"demo/module/note/notetransport/ginnote"
	"demo/module/upload/uploadtransport/ginupload"
	"demo/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx appctx.AppContext) {
	r.Use(middleware.Recover(appCtx)) // global middleware

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))

	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	v1.POST("/upload", ginupload.Upload(appCtx))

	notes := v1.Group("notes")
	{
		notes.POST("", middleware.RequiredAuth(appCtx), ginnote.CreateNote(appCtx))
		notes.GET("", ginnote.ListNote(appCtx))
		notes.PUT("/:note-id", middleware.RequiredAuth(appCtx), ginnote.UpdateNote(appCtx))
		notes.DELETE("/:note-id", middleware.RequiredAuth(appCtx), ginnote.DeleteNote(appCtx))
	}

	r.Static("/upload", "./static")

}

func setupAdminRouter(r *gin.Engine, appCtx appctx.AppContext) {

}
