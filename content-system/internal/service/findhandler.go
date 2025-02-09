package service

import (
	rpc "content-system/internal/api/content"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentFindReq struct {
	ID       int64  `json:"id"`        // 内容ID
	Author   string `json:"author"`    // 作者
	Title    string `json:"title"`     // 标题
	Page     int32  `json:"page"`      // 页
	PageSize int32  `json:"page_size"` // 页大小
}

func (c *CmsApp) ContentFind(ctx *gin.Context) {
	var req ContentFindReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsp, err := c.contentRpc.FindContent(ctx, &rpc.FindContentReq{
		Id:       req.ID,
		Author:   req.Author,
		Title:    req.Title,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
