package api

import (
	"content-system/internal/service"
	"github.com/gin-gonic/gin"
)

const (
	rootPath   = "/api/"
	noAuthPath = "/noauth/api/"
)

func NewRouters(r *gin.Engine) {
	app := service.NewCmsApp()
	session := NewSessionAuth()
	// 逻辑路由
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/hello
		//root.GET("/cms/hello", cmsApp.Hello)
		// /api/cms/create
		root.POST("/cms/content/create", app.ContentCreate)
		// /api/cms/update
		root.POST("/cms/content/update", app.ContentUpdate)
		// /api/cms/delete
		root.POST("/cms/content/delete", app.ContentDelete)
		// /api/cms/find
		root.POST("/cms/content/find", app.ContentFind)
	}
	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", app.Register)
		// /out/api/cms/login
		noAuth.POST("/cms/login", app.Login)
	}
}
