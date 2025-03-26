package idiom

import (
	"context"
	"encoding/json"
	"log"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/domain/websocket"
	"xiuyiPro/internal/service/live"
)

const (
	_giftExpireTime = 3600 * 24
)

func (s *Service) handleMsg(msg *websocket.Proto) (err error) {
	for index, cmd := range msg.BodyMuti {
		log.Printf("[WebsocketClient | msgResp] recv MsgResp index:%d ver:%d cmd:%s", index, msg.Version, string(cmd))
		switch string(cmd) {
		case live.OpenLiveCmdDm:
			// 弹幕
			data := new(live.OpenPlatformDM)
			if err = json.Unmarshal(msg.Body, data); err != nil {
				s.log.Error("handleMsg json.Unmarshal err(%+v)", err)
				return
			}
			if data.OpenId == "" || data.RoomId == 0 {
				s.log.Error("handleMsg data(%+v) invalid", data)
				return
			}
			return s.handleDm(context.Background(), data)
		case live.OpenLiveCmdGift:
			// 礼物
			data := new(live.OpenPlatformGift)
			if err = json.Unmarshal(msg.Body, data); err != nil {
				s.log.Error("handleMsg json.Unmarshal err(%+v)", err)
				return
			}
			if data.OpenId == "" || data.RoomId == 0 {
				s.log.Error("handleMsg data(%+v) invalid", data)
				return
			}
			return s.handleGift(context.Background(), data)
		}
	}
	return
}

func (s *Service) handleDm(ctx context.Context, data *live.OpenPlatformDM) (err error) {
	if data == nil || data.DmType != 1 || len(data.Msg) != 4 {
		return
	}
	// roomId -> gameId
	var (
		gameId string
	)
	if gameId, err = s.repo.GetRoomGame(ctx, data.RoomId); err != nil {
		s.log.WithContext(ctx).Errorf("handleDm s.repo.GetRoomGame err(%+v)", err)
		return
	}
	// 查看用户是否送过礼物
	exist, err := s.repo.GetGameGiftView(ctx, gameId, data.OpenId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("handleDm s.repo.GetGameGiftView err(%+v)", err)
		return
	}
	if !exist {
		s.log.WithContext(ctx).Info("handleDm user not send gift, data(%+v)", data)
		return
	}
	// 查看用户是否答对
	res, err := s.repo.GetGameIdiom(ctx, gameId)
	if err != nil || res == nil {
		s.log.WithContext(ctx).Errorf("handleDm s.repo.GetGameAnswer err(%+v)", err)
		return
	}
	if res.Answer != data.Msg {
		s.log.WithContext(ctx).Info("handleDm user answer wrong, data(%+v)", data)
		return
	}
	s.log.WithContext(ctx).Info("handleDm user answer right, data(%+v)", data)
	// 加分
	if _, err = s.UpdateRanking(ctx, &pb.UpdateRankingReq{
		GameId: gameId,
		Viewer: &pb.Viewer{
			Uid:  data.OpenId,
			Face: data.Uface,
			Name: data.Uname,
		},
	}); err != nil {
		s.log.WithContext(ctx).Errorf("handleDm s.UpdateRanking err(%+v)", err)
		return
	}
	return
}

func (s *Service) handleGift(ctx context.Context, data *live.OpenPlatformGift) (err error) {
	if data == nil || data.Price == 0 {
		return
	}
	// roomId -> gameId
	var (
		gameId string
	)
	if gameId, err = s.repo.GetRoomGame(ctx, data.RoomId); err != nil {
		s.log.WithContext(ctx).Errorf("handleGift s.repo.GetRoomGame err(%+v)", err)
		return
	}
	// 记录用户送礼物
	if err = s.repo.SetGameGiftView(ctx, gameId, data.OpenId, _giftExpireTime); err != nil {
		s.log.WithContext(ctx).Errorf("handleGift s.repo.SetGameGiftView err(%+v)", err)
		return
	}
	return
}
