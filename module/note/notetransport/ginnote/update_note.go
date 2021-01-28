package ginnote

import (
	"demo/common"
	"demo/module/note/notebusiness"
	"demo/module/note/notemodel"
	"demo/module/note/notestorge"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateNote(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.UpdateNote
		id, _ := strconv.Atoi(c.Param("note-id"))

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}

		data.Id = id
		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		biz := notebusiness.NewUpdateNoteBusiness(store)

		if err := biz.UpdateNote(c.Request.Context(), &data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
