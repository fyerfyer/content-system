package service

import (
	rpc "content-management/api/content"
	"content-management/internal/biz"
	"context"
	"time"
)

func (cs *ContentService) UpdateContent(ctx context.Context, req *rpc.UpdateContentReq) (*rpc.UpdateContentRsp, error) {
	content := req.GetContent()
	uc := cs.uc
	err := uc.UpdateContent(ctx, &biz.Content{
		ID:             content.Id,
		Title:          content.GetTitle(),
		VideoURL:       content.GetVideoUrl(),
		Author:         content.GetAuthor(),
		Description:    content.GetDescription(),
		Thumbnail:      content.GetThumbnail(),
		Category:       content.GetCategory(),
		Duration:       time.Duration(content.GetDuration()),
		Resolution:     content.GetResolution(),
		FileSize:       content.GetFileSize(),
		Format:         content.GetFormat(),
		Quality:        content.GetQuality(),
		ApprovalStatus: content.GetApprovalStatus(),
	})

	if err != nil {
		return nil, err
	}

	return &rpc.UpdateContentRsp{}, nil
}
