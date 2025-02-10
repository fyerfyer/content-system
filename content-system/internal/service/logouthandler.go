package service

import (
	"content-system/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

const SessionKey = "session_id"

type LogoutRsp struct {
	Message string `json:"message"`
}

func (c *CmsApp) Logout(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)
	if sessionID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	authKey := utils.GetAuthKey(sessionID)
	err := c.rdb.Del(ctx, authKey).Err()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to logout"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LogoutRsp{
			Message: "logout successful",
		},
	})
}
