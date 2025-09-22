package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

// BusinessRepo is a Greater repo.
type BusinessRepo interface {
	Reply(context.Context, *ReplyParam) (int64, error)
	AppealReview(context.Context, *AppealParam) (int64, error)
}

// BusinessUsecase is a Business usecase.
type BusinessUsecase struct {
	repo BusinessRepo
	log  *log.Helper
}

// NewBusinessUsecase new a Greeter usecase.
func NewBusinessUsecase(repo BusinessRepo, logger log.Logger) *BusinessUsecase {
	return &BusinessUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (b *BusinessUsecase) BizReply(ctx context.Context, param *ReplyParam) (int64, error) {
	// business reply user review
	b.log.WithContext(ctx).Infof("[biz] Reply(%+v)", param)
	return b.repo.Reply(ctx, param)

}

func (b *BusinessUsecase) BizAppealReview(ctx context.Context, param *AppealParam) (int64, error) {
	// business appeal user review
	b.log.WithContext(ctx).Infof("[biz] AppealReview(%+v)", param)
	return b.repo.AppealReview(ctx, param)
}
