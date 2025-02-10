package api

import (
	"content-system/internal/utils"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

const SessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{}
	connRdb(s)
	return s
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionID := ctx.GetHeader(SessionKey)
	if sessionID == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "session_id is null")
		return
	}
	authKey := utils.GetAuthKey(sessionID)
	loginTime, err := s.rdb.Get(ctx, authKey).Result()
	if err != nil && err != redis.Nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
		return
	}
	if loginTime == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "session auth fail")
		return
	}

	// Check if session is expired
	expirationTime, err := time.Parse(time.RFC3339, loginTime)
	if err != nil || time.Since(expirationTime) > 8*time.Hour {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, "session expired")
		return
	}

	ctx.Next()
}

func connRdb(s *SessionAuth) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	s.rdb = rdb
}
