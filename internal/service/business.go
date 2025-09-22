package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"review-b/internal/biz"

	pb "review-b/api/business/v1"
)

type BusinessService struct {
	pb.UnimplementedBusinessServer
	uc *biz.BusinessUsecase
}

func NewBusinessService(uc *biz.BusinessUsecase) *BusinessService {
	return &BusinessService{uc: uc}
}

func (s *BusinessService) ReplyUserReview(ctx context.Context, req *pb.ReplyReviewRequest) (*pb.ReplyReviewReply, error) {
	// 商家回复用户评价
	replyID, err := s.uc.BizReply(ctx, &biz.ReplyParam{
		ReviewID:  req.GetReviewID(),
		StoreID:   req.GetStoreID(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.ReplyReviewReply{ReplyID: replyID}, nil
}

func (s *BusinessService) AppealReview(ctx context.Context, req *pb.AppealReviewRequest) (*pb.AppealReviewReply, error) {
	log.Infof("[service] AppealReview,req:%v\n", req)
	// 商家申诉用户评价
	appealID, err := s.uc.BizAppealReview(ctx, &biz.AppealParam{
		ReviewID:  req.GetReviewID(),
		StoreID:   req.GetStoreID(),
		Reason:    req.GetReason(),
		Content:   req.GetContent(),
		PicInfo:   req.GetPicInfo(),
		VideoInfo: req.GetVideoInfo(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.AppealReviewReply{AppealID: appealID}, nil
}
