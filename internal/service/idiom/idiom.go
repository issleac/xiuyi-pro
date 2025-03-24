package idiom

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"sync"
	"time"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/biz"
	"xiuyiPro/internal/conf"
)

const _defaultRoomId = 1219

type Service struct {
	pb.UnimplementedIdiomServer
	repo  *biz.IdiomUsecase
	log   *log.Helper
	cfg   *conf.Application
	rooms sync.Map
}

func NewIdiomService(cfg *conf.Application, repo *biz.IdiomUsecase, logger log.Logger) *Service {
	s := &Service{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "service/idiom-service")),
		cfg:  cfg,
	}
	if cfg == nil {
		panic("config for NewTurtleService is nil")
	}
	return s
}

func (s *Service) SetIdiomBatch(ctx context.Context, req *pb.SetIdiomBatchReq) (resp *emptypb.Empty, err error) {
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

func (s *Service) GetIdiom(ctx context.Context, req *pb.GetIdiomReq) (resp *pb.GetIdiomResp, err error) {
	resp = new(pb.GetIdiomResp)
	if req == nil || req.RoomId == 0 {
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
	// 当前房间游戏绑定题目
	if err = s.repo.SetGameAnswer(ctx, req.RoomId, idiom.Answer, _gameExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("GetIdiom SetGameAnswer err(%+v)", err)
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

func (s *Service) generateID() string {
	return fmt.Sprintf("IDO-%d-%d", time.Now().UnixNano(), rand.Intn(10000))
}
