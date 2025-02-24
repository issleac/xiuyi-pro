package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"math/rand"
	"time"
	pb "xiuyiPro/api/turtle/v1"
	"xiuyiPro/internal/biz"
)

func (s *TurtleService) SetTurtleBatch(ctx context.Context, req *pb.SetTurtleBatchReq) (resp *pb.SetTurtleBatchResp, err error) {
	resp = new(pb.SetTurtleBatchResp)
	if req == nil || len(req.Turtles) == 0 {
		return nil, errors.BadRequest("参数错误", "SetTurtleBatch")
	}
	s.log.WithContext(ctx).Infof("SetTurtleBatch req(%+v)", req)
	var (
		turtles []*biz.Turtle
	)
	for _, t := range req.Turtles {
		turtles = append(turtles, &biz.Turtle{
			Qid:      s.generateQid(),
			Title:    t.Title,
			Content:  t.Content,
			Answer:   t.Answer,
			Category: int32(t.Category),
			Creator:  t.Creator,
			State:    int32(pb.Turtle_ONLINE),
		})
	}
	if _, err = s.repo.CreateTurtles(ctx, turtles); err != nil {
		s.log.WithContext(ctx).Errorf("CreateTurtles err(%+v)", err)
		return
	}
	return &pb.SetTurtleBatchResp{}, nil
}

func (s *TurtleService) GetTurtleList(ctx context.Context, req *pb.GetTurtleListReq) (resp *pb.GetTurtleListResp, err error) {
	resp = new(pb.GetTurtleListResp)
	if req == nil {
		return nil, errors.BadRequest("400", "GetTurtleList")
	}
	s.log.WithContext(ctx).Infof("GetTurtleList req(%+v)", req)
	var (
		limit, offset int32
		t             = new(biz.Turtle)
		res           []*biz.Turtle
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
	if res, err = s.repo.GetTurtleByPage(ctx, limit, offset, t); err != nil {
		s.log.WithContext(ctx).Errorf("GetTurtleByPage err(%+v)", err)
		return
	}
	for _, v := range res {
		resp.Turtles = append(resp.Turtles, &pb.Turtle{
			Qid:      v.Qid,
			Title:    v.Title,
			Content:  v.Content,
			Answer:   v.Answer,
			Category: pb.Turtle_Category(v.Category),
			Creator:  v.Creator,
			State:    pb.Turtle_State(v.State),
		})
	}
	return
}

// 生成唯一建qid
func (s *TurtleService) generateQid() string {
	// Turtle Soup
	return fmt.Sprintf("TS-%d-%d", time.Now().UnixNano(), rand.Intn(10000))
}
