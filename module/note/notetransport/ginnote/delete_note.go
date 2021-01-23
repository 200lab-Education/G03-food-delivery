package ginnote

import (
	"demo/common"
	"demo/module/note/notebusiness"
	"demo/module/note/notestorge"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteNote(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("note-id"))

		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		biz := notebusiness.NewDeleteNoteBiz(store)

		if err := biz.DeleteNote(id); err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
