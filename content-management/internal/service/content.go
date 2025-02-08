package service

import (
	content "content-management/api/content"
	"content-management/internal/biz"
)

// ContentService is a service for content.
type ContentService struct {
	content.UnsafeAppServer
	uc *biz.ContentUsecase
}

// NewContentService creates a new ContentService.
func NewContentService(uc *biz.ContentUsecase) *ContentService {
	return &ContentService{uc: uc}
}
