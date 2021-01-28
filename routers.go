package main

import (
	"demo/common"
	"demo/module/note/notetransport/ginnote"
	"github.com/gin-gonic/gin"
)

func setupRouter(r *gin.Engine, appCtx common.AppContext) {
	v1 := r.Group("/v1")
	notes := v1.Group("notes")
	{
		notes.PUT("/:note-id", ginnote.UpdateNote(appCtx))
		notes.POST("", ginnote.CreateNote(appCtx))
	}
}

func setupAdminRouter(r *gin.Engine, appCtx common.AppContext) {

}
