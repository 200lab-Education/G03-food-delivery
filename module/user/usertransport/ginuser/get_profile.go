package ginuser

import (
	"demo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfile(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		//data.Mask(true)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
