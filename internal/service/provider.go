package service

import (
	"github.com/whj1990/mine-grrpc/internal/repo"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewReviewService)

func NewReviewService(reviewProjectRepo repo.ReviewProjectRepo) ReviewService {
	return &reviewService{reviewProjectRepo}
}
