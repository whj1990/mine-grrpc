package service

import (
	"context"
	"mine/mine-grrpc/internal/repo"
	"mine/mine-grrpc/internal/service/stru"
	"mine/mine-grrpc/pbs"
)

type ReviewService interface {
	ReviewProjectList(context.Context, *pbs.ReviewProjectListParams) (*pbs.ReviewProjectListResponse, error)
}

type reviewService struct {
	reviewProjectRepo repo.ReviewProjectRepo
}

func (s *reviewService) ReviewProjectList(ctx context.Context, req *pbs.ReviewProjectListParams) (*pbs.ReviewProjectListResponse, error) {
	list, count, err := s.reviewProjectRepo.List(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := &pbs.ReviewProjectListResponse{
		Data: make([]*pbs.ReviewProjectData, 0),
	}
	for _, data := range *list {
		resp.Data = append(resp.Data, stru.Convert2ReviewProjectData(&data))
	}
	resp.PageNum = req.PageNum
	resp.PageSize = req.PageSize
	resp.Total = count
	return resp, err
}
