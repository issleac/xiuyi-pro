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
	if req.RoomId == 0 || req.Limit == 0 {
		s.log.WithContext(ctx).Warn("GetRanking req.RoomId is 0")
		return nil, errors.BadRequest("Invalid parameters", "GetRanking")
	}
	s.log.WithContext(ctx).Infof("GetRanking req(%+v)", req)
	var (
		gameId string
	)
	// room_id -> game_id
	if gameId, err = s.repo.GetRoomGame(ctx, req.RoomId); err != nil {
		s.log.WithContext(ctx).Errorf("GetIdiom GetRoomGame err(%+v)", err)
		return
	}
	// 从redis中获取排行榜
	res, err := s.repo.GetTopRanking(ctx, gameId, int64(req.Limit))
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
	if req.GetGameId() == "" || req.Viewer == nil || req.Viewer.Uid == "" {
		s.log.WithContext(ctx).Errorf("UpdateRanking req.GameId is nil, req(%+v)", req)
		return nil, errors.BadRequest("Invalid parameters", "UpdateRanking")
	}
	// 更新redis信息
	viewer := &biz.ViewerRanking{
		UID:  req.Viewer.Uid,
		Face: req.Viewer.Face,
		Name: req.Viewer.Name,
	}
	if err = s.repo.UpsertRanking(ctx, req.GetGameId(), viewer, _gameExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("UpdateRanking UpsertRanking gameId(%d)|viewer(%+v) err(%+v)", req.GetGameId(), viewer, err)
		return
	}
	return
}
