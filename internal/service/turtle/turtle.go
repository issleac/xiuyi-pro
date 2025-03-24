package turtle

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/types/known/emptypb"
	"math/rand"
	"time"
	pb "xiuyiPro/api/turtle/v1"
	"xiuyiPro/internal/biz"
	"xiuyiPro/internal/conf"
)

type Service struct {
	pb.UnimplementedTurtleServer
	repo *biz.TurtleUsecase
	log  *log.Helper
	cfg  *conf.Application
}

func NewTurtleService(cfg *conf.Application, repo *biz.TurtleUsecase, logger log.Logger) *Service {
	if cfg == nil {
		panic("config for NewTurtleService is nil")
	}
	s := &Service{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "service/turtle-service")),
		cfg:  cfg,
	}
	return s
}

func (s *Service) SetTurtleBatch(ctx context.Context, req *pb.SetTurtleBatchReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req == nil || len(req.Turtles) == 0 {
		return resp, errors.BadRequest("参数错误", "SetTurtleBatch")
	}
	s.log.WithContext(ctx).Infof("SetTurtleBatch req(%+v)", req)
	var (
		turtles []*biz.Turtle
	)
	for _, t := range req.Turtles {
		turtles = append(turtles, &biz.Turtle{
			Qid:        s.generateQid(),
			Title:      t.Title,
			Content:    t.Content,
			Answer:     t.Answer,
			Category:   int32(t.Category),
			Difficulty: int32(t.Difficulty),
			Creator:    t.Creator,
			State:      int32(pb.Turtle_ONLINE),
		})
	}
	if _, err = s.repo.CreateTurtles(ctx, turtles); err != nil {
		s.log.WithContext(ctx).Errorf("SetTurtleBatch CreateTurtles err(%+v)", err)
		return
	}
	return resp, nil
}

func (s *Service) GetTurtleList(ctx context.Context, req *pb.GetTurtleListReq) (resp *pb.GetTurtleListResp, err error) {
	resp = new(pb.GetTurtleListResp)
	if req == nil {
		return nil, errors.BadRequest("参数错误", "GetTurtleList")
	}
	s.log.WithContext(ctx).Infof("GetTurtleList req(%+v)", req)
	var (
		limit, offset int32
		t             = new(biz.Turtle)
		res           []*biz.Turtle
		total         int32
	)
	if req.GetPage() > 0 && req.GetPageSize() > 0 {
		limit = req.GetPageSize()
		offset = (req.GetPage() - 1) * req.GetPageSize()
	}
	if req.GetState() > 0 {
		t.State = int32(req.GetState())
	}
	if req.GetCategory() > 0 {
		t.Category = int32(req.GetCategory())
	}
	if req.GetDifficulty() > 0 {
		t.Difficulty = int32(req.GetDifficulty())
	}
	if res, total, err = s.repo.GetTurtleByPage(ctx, limit, offset, t); err != nil {
		s.log.WithContext(ctx).Errorf("GetTurtleByPage err(%+v)", err)
		return
	}
	for _, v := range res {
		resp.Turtles = append(resp.Turtles, &pb.Turtle{
			Qid:        v.Qid,
			Title:      v.Title,
			Content:    v.Content,
			Answer:     v.Answer,
			Category:   pb.Turtle_Category(v.Category),
			Difficulty: pb.Turtle_Degree(v.Difficulty),
			Creator:    v.Creator,
			State:      pb.Turtle_State(v.State),
		})
	}
	resp.Total = total
	return
}

// 生成唯一建qid
func (s *Service) generateQid() string {
	// Turtle Soup
	return fmt.Sprintf("TS-%d-%d", time.Now().UnixNano(), rand.Intn(10000))
}
