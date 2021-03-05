package ginnote

import (
	"demo/common"
	"demo/component/appctx"
	"demo/module/note/notebusiness"
	"demo/module/note/notemodel"
	"demo/module/note/notestorge"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateNote(provider appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data notemodel.UpdateNote
		//id, _ := strconv.Atoi(c.Param("note-id"))

		uid, err := common.FromBase58(c.Param("note-id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.Id = int(uid.GetLocalID())
		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		biz := notebusiness.NewUpdateNoteBusiness(store)

		if err := biz.UpdateNote(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
