package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
	"gvb_server/api/qq_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi  settings_api.SettingsApi
	ImageApi     images_api.ImagesApi
	AdvertiseApi advert_api.AdvertiseApi
	MenuApi      menu_api.MenuApi
	UserApi      user_api.UserApi
	TagApi       tag_api.TagApi
	MessageApi   message_api.MessageApi
	QQLoginApi   qq_api.QQLoginApi
}

var ApiGroupApp = new(ApiGroup)
