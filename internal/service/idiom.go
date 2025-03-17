package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"time"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/biz"
)

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
			Name:       i.Name,
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
		Name:       idiom.Name,
		Image:      idiom.Image,
		Difficulty: pb.Idiom_Degree(idiom.Difficulty),
		Creator:    idiom.Creator,
		State:      pb.Idiom_State(idiom.State),
	}
	return
}

func (s *IdiomService) generateID() string {
	return fmt.Sprintf("IDO-%d-%d", time.Now().UnixNano(), rand.Intn(10000))
}
