package service

import (
	"bytes"
	rpc "content-management/api/content"
	"content-management/internal/biz"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (cs *ContentService) CreateContent(ctx context.Context,
	req *rpc.CreateContentReq) (*rpc.CreateContentRsp, error) {
	content := req.GetContent()
	uc := cs.uc
	contentID := uuid.New().String()
	_, err := uc.CreateContent(ctx, &biz.Content{
		ContentID:      contentID,
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
	if err := cs.ExecFlow(contentID); err != nil {
		return nil, err
	}
	return &rpc.CreateContentRsp{}, nil
}

func (cs *ContentService) ExecFlow(contentID string) error {
	url := "http://localhost:7788/content-flow"
	method := "GET"
	payload := map[string]interface{}{
		"content_id": contentID,
	}
	data, _ := json.Marshal(payload)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return err
	}
	return nil
}
