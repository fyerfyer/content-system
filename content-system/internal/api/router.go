package api

import "github.com/gin-gonic/gin"

const (
	rootPath   = "/api/"
	noAuthPath = "/noauth/api/"
)

func NewRouters(r *gin.Engine) {
	cms
	// 逻辑路由
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/hello
		root.GET("/cms/hello", cmsApp.Hello)
		// /api/cms/create
		root.POST("/cms/content/create", cmsApp.ContentCreate)
		// /api/cms/update
		root.POST("/cms/content/update", cmsApp.ContentUpdate)
		// /api/cms/delete
		root.POST("/cms/content/delete", cmsApp.ContentDelete)
		// /api/cms/find
		root.POST("/cms/content/find", cmsApp.ContentFind)
	}
	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		// /out/api/cms/login
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}
