package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "xiuyiPro/api/turtle/v1"
	"xiuyiPro/internal/biz"
	"xiuyiPro/internal/conf"
	"xiuyiPro/internal/domain/live"
	"xiuyiPro/internal/domain/websocket"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewTurtleService)

type TurtleService struct {
	pb.UnimplementedTurtleServer
	repo       *biz.TurtleUsecase
	log        *log.Helper
	cfg        *conf.Application
	liveSingle *live.Live
	wsDao      *websocket.Websocket
}

func NewTurtleService(cfg *conf.Application, repo *biz.TurtleUsecase, logger log.Logger) *TurtleService {
	s := &TurtleService{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "service/turtle-service")),
		cfg:  cfg,
	}
	if cfg == nil {
		panic("config for NewTurtleService is nil")
	}
	var err error
	//s.liveSingle, err = live.New(cfg.GetLive(), s.log)
	//if err != nil {
	//	panic(err)
	//}
	s.wsDao, err = websocket.New(s.log)
	if err != nil {
		panic(err)
	}
	return s
}
