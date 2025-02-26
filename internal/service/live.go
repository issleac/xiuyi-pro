package service

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	errpb "xiuyiPro/api/helloworld/v1"
	pb "xiuyiPro/api/turtle/v1"
	liDm "xiuyiPro/internal/domain/live"
)

func (s *TurtleService) StartApp(ctx context.Context, req *pb.StartAppReq) (resp *pb.StartAppResp, err error) {
	resp = new(pb.StartAppResp)
	if req == nil {
		return resp, errors.BadRequest("参数错误", "StartApp")
	}
	s.log.WithContext(ctx).Infof("StartApp req(%+v)", req)

	// liveSingle 单例
	if s.liveSingle != nil {
		s.log.WithContext(ctx).Warn("StartApp liveSingle is not nil")
		err = errors.New(int(errpb.ErrorReason_WEBSOCKET_EXIST), errpb.ErrorReason_WEBSOCKET_EXIST.String(), "")
		return
	}
	if s.liveSingle, err = liDm.New(s.cfg.GetLive(), s.log); err != nil {
		s.log.WithContext(ctx).Errorf("StartApp live.New err(%+v)", err)
		return
	}

	// 开启应用
	sResp, err := s.liveSingle.StartApp(ctx, req.GetUpId(), s.cfg.GetLive().GetAppId())
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp s.liveDao.StartApp err(%+v)", err)
		return
	}
	// 解析返回值
	startAppRespData := &liDm.StartAppRespData{}
	err = json.Unmarshal(sResp.Data, &startAppRespData)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp json.Unmarshal err(%+v)", err)
		return
	}
	if startAppRespData == nil {
		err = errors.New(int(errpb.ErrorReason_START_APP_FAILED), errpb.ErrorReason_START_APP_FAILED.String(), "")
		s.log.WithContext(ctx).Error("StartApp startAppRespData is nil")
		return
	}

	if len(startAppRespData.WebsocketInfo.WssLink) == 0 {
		s.log.WithContext(ctx).Error("StartApp startAppRespData.WebsocketInfo.WssLink is nil")
		return
	}

	// todo，goroutine如何回收
	go func(ctx context.Context, gameId string) {
		for {
			time.Sleep(time.Second * 20)
			_, _ = s.liveSingle.AppHeart(ctx, gameId)
		}
	}(context.Background(), startAppRespData.GameInfo.GameId)

	// 开启长连
	err = s.liveSingle.WsDao.StartWebsocket(ctx, startAppRespData.WebsocketInfo.WssLink[0], startAppRespData.WebsocketInfo.AuthBody)
	if err != nil {
		if errors.Code(err) == int(errpb.ErrorReason_WEBSOCKET_EXIST) {
			s.log.WithContext(ctx).Warn("StartApp wsDao.StartWebsocket err(%+v)", err)
		} else {
			s.log.WithContext(ctx).Error("StartApp wsDao.StartWebsocket err(%+v)", err)
			return
		}
	}

	// 退出
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("WebsocketClient exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func (s *TurtleService) EndApp(ctx context.Context, req *pb.EndAppReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req == nil {
		return resp, errors.BadRequest("参数错误", "EndApp")
	}
	s.log.WithContext(ctx).Infof("EndApp req(%+v)", req)
	// todo：进程关闭，回收goroutine
	return
}
