package ginnote

import (
	"demo/common"
	"demo/component/appctx"
	"demo/module/note/notebusiness"
	"demo/module/note/notemodel"
	"demo/module/note/notestorge"
	"demo/module/upload/uploadstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateNote(provider appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.CreateNote

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		imgStore := uploadstorage.NewSQLStore(db)
		biz := notebusiness.NewCreateNoteBusiness(store, imgStore, provider.GetPubsub())

		if err := biz.CreateNote(c.Request.Context(), &data); err != nil {
			c.JSON(400, common.ErrInvalidRequest(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
		return
	}
}
