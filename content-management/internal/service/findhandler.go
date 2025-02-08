package service

import (
	rpc "content-management/api/content"
	"content-management/internal/biz"
	"context"
)

func (cs *ContentService) FindContent(ctx context.Context,
	req *rpc.FindContentReq) (*rpc.FindContentRsp, error) {
	findParams := &biz.FindParams{
		ID:       req.GetId(),
		Author:   req.GetAuthor(),
		Title:    req.GetTitle(),
		Page:     req.Page,
		PageSize: req.GetPageSize(),
	}
	uc := cs.uc
	results, total, err := uc.FindContent(ctx, findParams)
	if err != nil {
		return nil, err
	}
	var contents []*rpc.Content
	for _, r := range results {
		contents = append(contents, &rpc.Content{
			Id:             r.ID,
			Title:          r.Title,
			VideoUrl:       r.VideoURL,
			Author:         r.Author,
			Description:    r.Description,
			Thumbnail:      r.Thumbnail,
			Category:       r.Category,
			Duration:       r.Duration.Milliseconds(),
			Resolution:     r.Resolution,
			FileSize:       r.FileSize,
			Format:         r.Format,
			Quality:        r.Quality,
			ApprovalStatus: r.ApprovalStatus,
		})
	}
	rsp := &rpc.FindContentRsp{
		Total:    total,
		Contents: contents,
	}
	return rsp, nil
}
