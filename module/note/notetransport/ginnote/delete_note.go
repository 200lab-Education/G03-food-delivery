package ginnote

import (
	"demo/common"
	"demo/component/appctx"
	"demo/module/note/notebusiness"
	"demo/module/note/notestorge"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteNote(provider appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("note-id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		biz := notebusiness.NewDeleteNoteBiz(store, requester)

		if err := biz.DeleteNote(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
