package ginuser

import (
	"demo/common"
	"demo/component/appctx"
	"demo/component/hasher"
	"demo/component/tokenprovider/jwt"
	"demo/module/user/userbusiness"
	"demo/module/user/usermodel"
	"demo/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		business := userbusiness.NewLoginBusiness(store, tokenProvider, md5, appctx.NewTokenConfig())
		account, err := business.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
