package appctx

import (
	"demo/component/uploadprovider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
}

type ctx struct {
	mainDB     *gorm.DB
	secret     string
	upProvider uploadprovider.UploadProvider
}

func New(mainDB *gorm.DB, secret string, upProvider uploadprovider.UploadProvider) *ctx {
	return &ctx{mainDB: mainDB, secret: secret, upProvider: upProvider}
}

func (c ctx) GetMainDBConnection() *gorm.DB {
	return c.mainDB
}

func (c ctx) SecretKey() string {
	return c.secret
}

func (c ctx) UploadProvider() uploadprovider.UploadProvider {
	return c.upProvider
}

type tokenExpiry struct {
	atExp int
	rtExp int
}

func NewTokenConfig() tokenExpiry {
	return tokenExpiry{
		atExp: 60 * 60 * 24 * 7,
		rtExp: 60 * 60 * 24 * 7 * 2,
	}
}

func (tk tokenExpiry) GetAtExp() int {
	return tk.atExp
}

func (tk tokenExpiry) GetRtExp() int {
	return tk.rtExp
}
