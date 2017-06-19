package auth

import (
	"github.com/qor/auth"
	"github.com/qor/auth/database"
	"github.com/qor/auth/oauth/twitter"
	"github.com/qor/auth/phone"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/config"
	"github.com/qor/qor-example/db"
)

var Auth = auth.New(&auth.Config{
	DB:        db.DB,
	Render:    config.View,
	UserModel: models.User{},
})

func init() {
	Auth.RegisterProvider(database.New())
	Auth.RegisterProvider(phone.New())
	Auth.RegisterProvider(twitter.New(&config.Config.Twitter))
}
