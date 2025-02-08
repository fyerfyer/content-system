package service

import (
	rpc "content-management/api/content"
	"context"
)

func (cs *ContentService) DeleteContent(ctx context.Context,
	req *rpc.DeleteContentReq) (*rpc.DeleteContentRsp, error) {
	uc := cs.uc
	err := uc.DeleteContent(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &rpc.DeleteContentRsp{}, nil
}
