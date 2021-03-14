package appctx

import (
	"demo/component/uploadprovider"
	"demo/pubsub"
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	UploadProvider() uploadprovider.UploadProvider
	GetPubsub() pubsub.Pubsub
}

type ctx struct {
	mainDB     *gorm.DB
	secret     string
	upProvider uploadprovider.UploadProvider
	ps         pubsub.Pubsub
}

func New(mainDB *gorm.DB, secret string, upProvider uploadprovider.UploadProvider, ps pubsub.Pubsub) *ctx {
	return &ctx{mainDB: mainDB, secret: secret, upProvider: upProvider, ps: ps}
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

func (c ctx) GetPubsub() pubsub.Pubsub {
	return c.ps
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
