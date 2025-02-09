package service

import (
	rpc "content-system/internal/api/content"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContentDeleteReq struct {
	ID int64 `json:"id" binding:"required"` // 内容ID
}

type ContentDeleteRsp struct {
	Message string `json:"message"`
}

func (c *CmsApp) ContentDelete(ctx *gin.Context) {
	var req ContentDeleteReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsp, err := c.contentRpc.DeleteContent(ctx, &rpc.DeleteContentReq{
		Id: req.ID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
