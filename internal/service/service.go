package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "xiuyiPro/api/turtle/v1"
	"xiuyiPro/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService, NewTurtleService)

type TurtleService struct {
	pb.UnimplementedTurtleServer
	repo *biz.TurtleUsecase
	log  *log.Helper
}

func NewTurtleService(repo *biz.TurtleUsecase, logger log.Logger) *TurtleService {
	return &TurtleService{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "service/turtle-service"))}
}
