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
	if req == nil || len(req.GetUpCodeId()) == 0 || req.GetRoomId() <= 0 {
		return resp, errors.BadRequest("参数错误", "StartApp") //todo：错误码
	}
	s.log.WithContext(ctx).Infof("StartApp req(%+v)", req)
	var (
		startAppRespData = new(live.StartAppRespData)
		key, secret      = s.cfg.GetLive().GetAccessKey(), s.cfg.GetLive().GetAccessSecret()
		appId            = s.cfg.GetLive().AppId
		host             = s.cfg.GetLive().GetHost()
	)
	// 开启app
	sResp, err := live.StartApp(ctx, key, secret, host, req.UpCodeId, appId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp s.liveDao.StartApp key(%s)|secret(%s) err(%+v)", key, secret, err)
		return
	}
	// 解析返回值
	err = json.Unmarshal(sResp.Data, startAppRespData)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp json.Unmarshal err(%+v)", err)
		return
	}
	if startAppRespData == nil {
		err = errors.New(500, "StartApp startAppRespData is nil", "") //todo：错误码
		s.log.WithContext(ctx).Error("StartApp startAppRespData is nil")
		return
	}
	if len(startAppRespData.WebsocketInfo.WssLink) == 0 {
		err = errors.New(500, "StartApp startAppRespData.WebsocketInfo.WssLink is nil", "") //todo：错误码
		s.log.WithContext(ctx).Error("StartApp startAppRespData.WebsocketInfo.WssLink is nil")
		return
	}

	// 保持心跳
	// todo：暂停心跳，回收goroutine
	go func(ctx context.Context, gameId string) {
		for {
			time.Sleep(time.Second * 20)
			_, _ = live.AppHeart(ctx, key, secret, host, gameId)
		}
	}(context.Background(), startAppRespData.GameInfo.GameId)

	// 开启长连
	ws, err := websocket.New(s.log)
	if err != nil {
		s.log.WithContext(ctx).Error("StartApp websocket.New err(%+v)", err)
		return
	}
	err = ws.StartWebsocket(ctx, startAppRespData.WebsocketInfo.WssLink[0], startAppRespData.WebsocketInfo.AuthBody, s.handleMsg)
	if err != nil {
		s.log.WithContext(ctx).Error("StartApp ws.StartWebsocket err(%+v)", err)
		return
	}

	// 保存房间信息
	if err = s.repo.SetRoomGame(ctx, req.RoomId, startAppRespData.GameInfo.GameId, _gameExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("StartApp s.repo.SetRoomGame err(%+v)", err)
		return
	}
	resp.GameId = startAppRespData.GameInfo.GameId
	return
}

func (s *Service) EndApp(ctx context.Context, req *pb.EndAppReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req == nil {
		return resp, errors.BadRequest("参数错误", "EndApp") //todo：错误码
	}
	s.log.WithContext(ctx).Infof("EndApp req(%+v)", req)
	// room_id -> game_id
	gameId, err := s.repo.GetRoomGame(ctx, req.GetRoomId())
	if err != nil {
		s.log.WithContext(ctx).Errorf("EndApp s.repo.GetRoomGame err(%+v)", err)
		return
	}
	var (
		key, secret = s.cfg.GetLive().GetAccessKey(), s.cfg.GetLive().GetAccessSecret()
		appId       = s.cfg.GetLive().AppId
		host        = s.cfg.GetLive().GetHost()
	)
	_, err = live.EndApp(ctx, key, secret, host, gameId, appId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("EndApp s.liveDao.EndApp gameId(%s)|appId(%d) err(%+v)", gameId, appId, err)
		return nil, err
	}
	// 删除房间信息
	if err = s.repo.DelRoomGame(ctx, req.GetRoomId()); err != nil {
		s.log.WithContext(ctx).Errorf("EndApp s.repo.DelRoomGame err(%+v)", err)
		return
	}
	// todo：进程关闭，回收goroutine
	return
}
