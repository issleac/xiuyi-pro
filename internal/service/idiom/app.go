package idiom

import (
	"context"
	"encoding/json"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/errorcode"
	"xiuyiPro/internal/domain/websocket"
	"xiuyiPro/internal/service/live"
)

func (s *Service) StartApp(ctx context.Context, req *pb.StartAppReq) (resp *pb.StartAppResp, err error) {
	resp = new(pb.StartAppResp)
	if req == nil || len(req.GetUpCodeId()) == 0 || req.GetRoomId() <= 0 {
		return resp, errorcode.ParamInvalid
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
		err = errorcode.StartAppFailed
		return
	}
	// 解析返回值
	err = json.Unmarshal(sResp.Data, startAppRespData)
	if err != nil {
		s.log.WithContext(ctx).Errorf("StartApp json.Unmarshal err(%+v)", err)
		err = errorcode.StartAppFailed
		return
	}
	if startAppRespData == nil {
		s.log.WithContext(ctx).Error("StartApp startAppRespData is nil")
		err = errorcode.StartAppFailed
		return
	}
	if len(startAppRespData.WebsocketInfo.WssLink) == 0 {
		s.log.WithContext(ctx).Error("StartApp startAppRespData.WebsocketInfo.WssLink is nil")
		err = errorcode.StartAppFailed
		return
	}

	// 保持心跳
	gCtx, cancel := context.WithCancel(context.Background())
	if err = s.repo.SetGameCancelFunc(gCtx, startAppRespData.GameInfo.GameId, cancel, _gameExpireTime); err != nil {
		s.log.WithContext(gCtx).Errorf("StartApp s.repo.SetGameCancelFunc err(%+v)", err)
		err = errorcode.ServiceError
		return
	}
	go func(gCtx context.Context, gameId string) {
		for {
			select {
			case <-gCtx.Done():
				return
			case <-time.After(time.Second * 20):
				_, _ = live.AppHeart(gCtx, key, secret, host, gameId)
			}
		}
	}(gCtx, startAppRespData.GameInfo.GameId)

	// 开启长连
	ws, err := websocket.New(s.log)
	if err != nil {
		s.log.WithContext(ctx).Error("StartApp websocket.New err(%+v)", err)
		err = errorcode.CreateWsFailed
		return
	}
	err = ws.StartWebsocket(ctx, gCtx, startAppRespData.WebsocketInfo.WssLink[0], startAppRespData.WebsocketInfo.AuthBody, s.handleMsg)
	if err != nil {
		s.log.WithContext(ctx).Error("StartApp ws.StartWebsocket err(%+v)", err)
		err = errorcode.CreateWsFailed
		return
	}

	// 保存房间信息
	if err = s.repo.SetRoomGame(ctx, req.RoomId, startAppRespData.GameInfo.GameId, _gameExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("StartApp s.repo.SetRoomGame err(%+v)", err)
		err = errorcode.ServiceError
		return
	}
	resp.GameId = startAppRespData.GameInfo.GameId
	return
}

func (s *Service) EndApp(ctx context.Context, req *pb.EndAppReq) (resp *emptypb.Empty, err error) {
	resp = new(emptypb.Empty)
	if req == nil {
		return nil, errorcode.ParamInvalid
	}
	s.log.WithContext(ctx).Infof("EndApp req(%+v)", req)
	// room_id -> game_id
	gameId, err := s.repo.GetRoomGame(ctx, req.GetRoomId())
	if err != nil {
		s.log.WithContext(ctx).Errorf("EndApp s.repo.GetRoomGame err(%+v)", err)
		err = errorcode.ServiceError
		return
	}
	if gameId == "" {
		s.log.WithContext(ctx).Errorf("EndApp s.repo.GetRoomGame gameId(%s) err(%+v)", gameId, err)
		err = errorcode.GameNotFound
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
		err = errorcode.CloseAppFailed
		return
	}
	// 删除房间信息
	if err = s.repo.DelRoomGame(ctx, req.GetRoomId()); err != nil {
		s.log.WithContext(ctx).Errorf("EndApp s.repo.DelRoomGame err(%+v)", err)
		err = errorcode.ServiceError
		return
	}
	// 回收goroutine
	cancel, err := s.repo.GetGameCancelFunc(ctx, gameId)
	if err != nil || cancel == nil {
		s.log.WithContext(ctx).Errorf("EndApp s.repo.GetGameCancelFunc err(%+v)", err)
		err = errorcode.ServiceError
		return
	}
	cancel()
	return
}
