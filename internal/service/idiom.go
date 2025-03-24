package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"strconv"
	"time"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/biz"
)

const _defaultRoomId = 1219

func (s *IdiomService) SetIdiomBatch(ctx context.Context, req *pb.SetIdiomBatchReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req == nil || len(req.Idioms) == 0 {
		return resp, errors.BadRequest("Invalid parameters", "SetIdiomBatch")
	}
	s.log.WithContext(ctx).Infof("SetIdiomBatch req(%+v)", req)
	var (
		idioms []*biz.Idiom
	)
	for _, i := range req.Idioms {
		idioms = append(idioms, &biz.Idiom{
			Iid:        s.generateID(),
			Answer:     i.Answer,
			Image:      i.Image,
			Difficulty: int32(i.Difficulty),
			Creator:    i.Creator,
			State:      int32(pb.Idiom_ONLINE),
		})
	}
	if _, err = s.repo.CreateIdioms(ctx, idioms); err != nil {
		s.log.WithContext(ctx).Errorf("SetIdiomBatch SaveBatch err(%+v)", err)
		return
	}
	return resp, nil
}

func (s *IdiomService) GetIdiom(ctx context.Context, req *pb.GetIdiomReq) (resp *pb.GetIdiomResp, err error) {
	resp = new(pb.GetIdiomResp)
	if req == nil {
		return nil, errors.BadRequest("Invalid parameters", "GetIdiomList")
	}
	s.log.WithContext(ctx).Infof("GetIdiom req(%+v)", req)
	var (
		idiom *biz.Idiom
	)
	if idiom, err = s.repo.FindByID(ctx, req.Id); err != nil {
		s.log.WithContext(ctx).Errorf("GetIdiom FindByID err(%+v)", err) //todo:404
		return
	}
	resp.Idiom = &pb.Idiom{
		Iid:        idiom.Iid,
		Answer:     idiom.Answer,
		Image:      idiom.Image,
		Difficulty: pb.Idiom_Degree(idiom.Difficulty),
		Creator:    idiom.Creator,
		State:      pb.Idiom_State(idiom.State),
	}
	return
}

// GetRanking 获取排行榜
func (s *IdiomService) GetRanking(ctx context.Context, req *pb.GetRankingReq) (resp *pb.GetRankingResp, err error) {
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
	// todo:获取face和name
	for _, r := range res {
		resp.Viewers = append(resp.Viewers, &pb.Viewer{
			Uid:    r.UID,
			Face:   "",
			Name:   "",
			Source: r.Score,
			Index:  r.Index,
		})
	}
	return
}

// UpdateRanking 更新排行榜
func (s *IdiomService) UpdateRanking(ctx context.Context, req *pb.UpdateRankingReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req.RoomId == 0 {
		s.log.Warn("GetRanking req.RoomId is 0")
		req.RoomId = _defaultRoomId
	}
	if req.Uid <= 0 {
		return nil, errors.BadRequest("Invalid parameters", "UpdateRanking")
	}
	// 更新redis信息
	if err = s.repo.UpsertRanking(ctx, req.RoomId, strconv.FormatInt(req.Uid, 10)); err != nil {
		s.log.WithContext(ctx).Errorf("UpdateRanking UpsertRanking err(%+v)", err)
		return
	}
	return
}

func (s *IdiomService) generateID() string {
	return fmt.Sprintf("IDO-%d-%d", time.Now().UnixNano(), rand.Intn(10000))
}
