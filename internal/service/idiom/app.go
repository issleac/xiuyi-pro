package idiom

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/domain/websocket"
	"xiuyiPro/internal/service/live"
)

func (s *Service) StartApp(ctx context.Context, req *pb.StartAppReq) (resp *pb.StartAppResp, err error) {
	resp = new(pb.StartAppResp)
	if req == nil || len(req.GetUpCodeId()) == 0 {
		return resp, errors.BadRequest("参数错误", "StartApp")
	}
	s.log.WithContext(ctx).Infof("StartApp req(%+v)", req)
	var (
		gameId string
		game   = &live.RoomGame{
			RoomId:   req.GetRoomId(),
			UpCodeId: req.GetUpCodeId(),
		}
		startAppRespData = new(live.StartAppRespData)
		key, secret      = s.cfg.GetLive().GetAccessKey(), s.cfg.GetLive().GetAccessSecret()
		appId            = s.cfg.GetLive().AppId
		host             = s.cfg.GetLive().GetHost()
	)
	sResp, err := live.StartApp(ctx, key, secret, host, req.UpCodeId, appId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp s.liveDao.StartApp err(%+v)", err)
		return
	}
	// 解析返回值
	err = json.Unmarshal(sResp.Data, &startAppRespData)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp json.Unmarshal err(%+v)", err)
		return
	}
	if startAppRespData == nil {
		err = errors.New(500, "StartApp startAppRespData is nil", "")
		s.log.WithContext(ctx).Error("StartApp startAppRespData is nil")
		return
	}
	if len(startAppRespData.WebsocketInfo.WssLink) == 0 {
		s.log.WithContext(ctx).Error("StartApp startAppRespData.WebsocketInfo.WssLink is nil")
		return
	}

	go func(ctx context.Context, gameId string) {
		for {
			time.Sleep(time.Second * 20)
			_, _ = live.AppHeart(ctx, key, secret, host, gameId)
		}
	}(context.Background(), startAppRespData.GameInfo.GameId)

	// 开启长连
	ws, _ := websocket.New(s.log)
	err = ws.StartWebsocket(ctx, startAppRespData.WebsocketInfo.WssLink[0], startAppRespData.WebsocketInfo.AuthBody, s.handleMsg)
	if err != nil {
		s.log.WithContext(ctx).Error("StartApp ws.StartWebsocket err(%+v)", err)
		return
	}
	game.Ws = ws

	s.rooms.Store(gameId, game)
	return
}

func (s *Service) EndApp(ctx context.Context, req *pb.EndAppReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req == nil {
		return resp, errors.BadRequest("参数错误", "EndApp")
	}
	s.log.WithContext(ctx).Infof("EndApp req(%+v)", req)
	if _, ok := s.rooms.Load(req.GetGameId()); !ok {
		err = errors.BadRequest("房间不存在", "EndApp")
		return
	}
	s.rooms.Delete(req.GetGameId())
	// todo：进程关闭，回收goroutine
	return
}
