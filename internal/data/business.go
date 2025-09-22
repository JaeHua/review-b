package data

import (
	"context"
	v1 "review-b/api/review/v1"

	"review-b/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type businessRepo struct {
	data *Data
	log  *log.Helper
}

// NewBusinessRepo .
func NewBusinessRepo(data *Data, logger log.Logger) biz.BusinessRepo {
	return &businessRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *businessRepo) Reply(ctx context.Context, param *biz.ReplyParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] Reply param(%+v)", param)
	// 之前都是写数据库
	// 现在需要通过 RPC 调用 review-service 服务的接口
	ret, err := r.data.rc.ReplyReview(ctx, &v1.ReplyReviewRequest{
		ReviewID:  param.ReviewID,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[data] Reply call review-service ReplyReview error(%v)", err)
		return 0, err
	}
	return ret.ReplyID, nil
}

func (r *businessRepo) AppealReview(ctx context.Context, param *biz.AppealParam) (int64, error) {
	r.log.WithContext(ctx).Infof("[data] AppealReview param(%+v)", param)
	// 之前都是写数据库
	// 现在需要通过 RPC 调用 review-service 服务的接口
	ret, err := r.data.rc.AppealReview(ctx, &v1.AppealReviewRequest{
		ReviewID:  param.ReviewID,
		Reason:    param.Reason,
		StoreID:   param.StoreID,
		Content:   param.Content,
		PicInfo:   param.PicInfo,
		VideoInfo: param.VideoInfo,
	})
	if err != nil {
		r.log.WithContext(ctx).Errorf("[data] AppealReview call review-service AppealReview error(%v)", err)
		return 0, err
	}
	return ret.AppealID, nil
}
