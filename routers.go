package main

import (
	"demo/common"
	"demo/middleware"
	"demo/module/note/notetransport/ginnote"
	"demo/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx common.AppContext) {
	r.Use(middleware.Recover(appCtx)) // global middleware

	v1 := r.Group("/v1")

	v1.POST("/register", ginuser.Register(appCtx))

	notes := v1.Group("notes")
	{
		notes.POST("", ginnote.CreateNote(appCtx))
		notes.GET("", ginnote.ListNote(appCtx))
		notes.PUT("/:note-id", ginnote.UpdateNote(appCtx))
	}

}

func setupAdminRouter(r *gin.Engine, appCtx common.AppContext) {

}
