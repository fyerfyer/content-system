package service

import (
	rpc "content-system/internal/api/content"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ContentUpdateReq struct {
	ID             int64         `json:"id" binding:"required"` // 内容标题
	Title          string        `json:"title"`                 // 内容标题
	VideoURL       string        `json:"video_url"`             // 视频播放URL
	Author         string        `json:"author"`                // 作者
	Description    string        `json:"description"`           // 内容描述
	Thumbnail      string        `json:"thumbnail"`             // 封面图URL
	Category       string        `json:"category"`              // 内容分类
	Duration       time.Duration `json:"duration"`              // 内容时长
	Resolution     string        `json:"resolution"`            // 分辨率 如720p、1080p
	FileSize       int64         `json:"fileSize"`              // 文件大小
	Format         string        `json:"format"`                // 文件格式 如MP4、AVI
	Quality        int32         `json:"quality"`               // 视频质量 1-高清 2-标清
	ApprovalStatus int32         `json:"approval_status"`       // 审核状态 1-审核中 2-审核通过 3-审核不通过
}

type ContentUpdateRsp struct {
	Message string `json:"message"`
}

func (c *CmsApp) ContentUpdate(ctx *gin.Context) {
	var req ContentUpdateReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	rsp, err := c.contentRpc.UpdateContent(ctx, &rpc.UpdateContentReq{
		Content: &rpc.Content{
			Id:             req.ID,
			Title:          req.Title,
			Description:    req.Description,
			Author:         req.Author,
			VideoUrl:       req.VideoURL,
			Thumbnail:      req.Thumbnail,
			Category:       req.Category,
			Duration:       req.Duration.Milliseconds(),
			Resolution:     req.Resolution,
			FileSize:       req.FileSize,
			Format:         req.Format,
			Quality:        req.Quality,
			ApprovalStatus: req.ApprovalStatus,
		},
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
