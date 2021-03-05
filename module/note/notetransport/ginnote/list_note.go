package ginnote

import (
	"demo/common"
	"demo/component/appctx"
	"demo/module/note/notebusiness"
	"demo/module/note/notestorge"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func demoCrash() {
	var a []int
	fmt.Println(a[0])
}

func ListNote(provider appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		//demoCrash()
		//id, _ := strconv.Atoi(c.Param("note-id"))
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := provider.GetMainDBConnection()
		store := notestorge.NewSQLStore(db)
		biz := notebusiness.NewListNoteBiz(store)

		result, err := biz.ListNote(c.Request.Context(), &paging)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()

			if i == len(result)-1 {
				paging.NextCursor = result[i].FakeId.String()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}
