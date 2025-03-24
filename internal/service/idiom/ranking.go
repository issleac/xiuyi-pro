package idiom

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/biz"
)

const (
	// 24小时
	_gameExpireTime = 3600 * 24
)

// GetRanking 获取排行榜
func (s *Service) GetRanking(ctx context.Context, req *pb.GetRankingReq) (resp *pb.GetRankingResp, err error) {
	resp = new(pb.GetRankingResp)
	if req.RoomId == 0 {
		s.log.Warn("GetRanking req.RoomId is 0")
		req.RoomId = _defaultRoomId
	}
	if req.Limit == 0 {
		s.log.Error("GetRanking req.Limit is 0")
		return nil, errors.BadRequest("Invalid parameters", "GetRanking")
	}
	s.log.WithContext(ctx).Infof("GetRanking req(%+v)", req)
	// 从redis中获取排行榜
	res, err := s.repo.GetTopRanking(ctx, req.RoomId, int64(req.Limit))
	if err != nil {
		s.log.WithContext(ctx).Errorf("GetRanking GetTopRanking err(%+v)", err)
		return
	}
	for _, r := range res {
		resp.Viewers = append(resp.Viewers, &pb.Viewer{
			Uid:    r.UID,
			Face:   r.Face,
			Name:   r.Name,
			Source: r.Score,
			Index:  r.Index,
		})
	}
	return
}

// UpdateRanking 更新排行榜
func (s *Service) UpdateRanking(ctx context.Context, req *pb.UpdateRankingReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req.RoomId == 0 || req.Viewer == nil {
		s.log.Warn("GetRanking req(%+v) invalid", req)
		req.RoomId = _defaultRoomId
	}
	if req.Viewer.Uid == "" {
		return nil, errors.BadRequest("Invalid parameters", "UpdateRanking")
	}
	// 更新redis信息
	viewer := &biz.ViewerRanking{
		UID:  req.Viewer.Uid,
		Face: req.Viewer.Face,
		Name: req.Viewer.Name,
	}
	if err = s.repo.UpsertRanking(ctx, req.RoomId, viewer, _gameExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("UpdateRanking UpsertRanking err(%+v)", err)
		return
	}
	return
}
