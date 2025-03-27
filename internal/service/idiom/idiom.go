package idiom

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"time"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/errorcode"
	"xiuyiPro/internal/biz"
	"xiuyiPro/internal/conf"
)

const _defaultRoomId = 1219

type Service struct {
	pb.UnimplementedIdiomServer
	repo *biz.IdiomUsecase
	log  *log.Helper
	cfg  *conf.Application
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
		return resp, errorcode.ParamInvalid
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
		err = errorcode.ServiceError
		return
	}
	return resp, nil
}

func (s *Service) GetIdiom(ctx context.Context, req *pb.GetIdiomReq) (resp *pb.GetIdiomResp, err error) {
	resp = new(pb.GetIdiomResp)
	if req == nil || req.RoomId == 0 {
		return nil, errorcode.ParamInvalid
	}
	s.log.WithContext(ctx).Infof("GetIdiom req(%+v)", req)
	var (
		idiom     *biz.Idiom
		gameId    string
		lastIdiom *biz.IdiomAnswer
	)
	// roomId -> gameId
	if gameId, err = s.repo.GetRoomGame(ctx, req.RoomId); err != nil {
		s.log.WithContext(ctx).Errorf("GetIdiom GetRoomGame err(%+v)", err)
		err = errorcode.ServiceError
		return
	}
	if gameId == "" {
		s.log.WithContext(ctx).Errorf("GetIdiom GetRoomGame gameId(%s) err(%+v)", gameId, err)
		err = errorcode.GameNotFound
		return
	}
	// 获取下一个题目
	if req.Id == 0 {
		// 获取当前房间游戏绑定题目
		lastIdiom, err = s.repo.GetGameIdiom(ctx, gameId)
		if err != nil {
			s.log.WithContext(ctx).Errorf("GetIdiom GetGameIdiom err(%+v)", err)
			err = errorcode.ServiceError
			return
		}
		if lastIdiom == nil {
			req.Id = 1
		} else {
			req.Id = lastIdiom.IdiomId + 1
		}
		s.log.WithContext(ctx).Infof("GetIdiom GetGameIdiom lastIdiom(%+v) req.Id(%+v)", lastIdiom, req.Id)
	}
	if idiom, err = s.repo.FindByGTEID(ctx, req.Id); err != nil {
		s.log.WithContext(ctx).Errorf("GetIdiom FindByGTEID req.Id(%d) err(%+v)", req.Id, err)
		return
	}
	if idiom == nil {
		// 循环找题目
		req.Id = 1
		if idiom, err = s.repo.FindByGTEID(ctx, req.Id); err != nil {
			s.log.WithContext(ctx).Errorf("GetIdiom FindByGTEID req.Id(%d) err(%+v)", req.Id, err)
			err = errorcode.ServiceError
			return
		}
		if idiom == nil {
			err = errorcode.NoAvailableIdiom
			return
		}
		return
	}
	// 更新当前房间游戏绑定题目
	if err = s.repo.SetGameIdiom(ctx, gameId, idiom.ID, idiom.Answer, _gameExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("GetIdiom SetGameAnswer idiom(%+v) err(%+v)", idiom, err)
		err = errorcode.ServiceError
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
