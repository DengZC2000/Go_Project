package routers

import (
	"github.com/gin-gonic/gin"
	"gvb_server/api"
)

func MenuRouter(enter *gin.RouterGroup) {
	MenuApi := api.ApiGroupApp.MenuApi
	enter.POST("/menu", MenuApi.MenuCreateView)
	enter.GET("/menu", MenuApi.MenuListView)
	enter.GET("/menu_names", MenuApi.MenuNameListView)
	enter.PUT("/menu/:id", MenuApi.MenuUpdateView)
	enter.DELETE("/menu", MenuApi.MenuDeleteView)
	enter.POST("/menu/:id", MenuApi.MenuDetailView)
}
