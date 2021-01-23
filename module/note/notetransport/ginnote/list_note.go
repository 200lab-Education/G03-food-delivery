package ginnote

import (
	"demo/common"
	"demo/module/note/notebusiness"
	"demo/module/note/notestorge"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListNote(provider common.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		//id, _ := strconv.Atoi(c.Param("note-id"))
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}

		paging.Fulfill()

		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		biz := notebusiness.NewListNoteBiz(store)

		result, err := biz.ListNote(&paging)

		if err != nil {
			c.JSON(401, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
