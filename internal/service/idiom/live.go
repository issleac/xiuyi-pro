package idiom

import (
	"context"
	"encoding/json"
	"log"
	pb "xiuyiPro/api/idiom/v1"
	"xiuyiPro/internal/domain/websocket"
	"xiuyiPro/internal/service/live"
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
		}
	}
	return
}

func (s *Service) handleDm(ctx context.Context, data *live.OpenPlatformDM) (err error) {
	if data == nil || data.DmType != 1 || len(data.Msg) != 4 {
		return
	}
	// 查看用户是否答对
	res, err := s.repo.GetGameAnswer(ctx, data.RoomId)
	if err != nil {
		s.log.WithContext(ctx).Errorf("handleDm s.repo.GetGameAnswer err(%+v)", err)
		return
	}
	if res != data.Msg {
		s.log.WithContext(ctx).Info("handleDm user answer wrong, data(%+v)", data)
		return
	}
	s.log.WithContext(ctx).Info("handleDm user answer right, data(%+v)", data)
	// 加分
	if _, err = s.UpdateRanking(ctx, &pb.UpdateRankingReq{
		RoomId: data.RoomId,
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
