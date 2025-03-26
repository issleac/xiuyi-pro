package live

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	_startPath = "/v2/app/start"
	_heartPath = "/v2/app/heartbeat"
	_endPath   = "/v2/app/end"
)

// StartApp 启动app
func StartApp(c context.Context, key, secret, host, code string, appId int64) (resp BaseResp, err error) {
	startAppReq := StartAppRequest{
		Code:  code,
		AppId: appId,
	}
	reqJson, err := json.Marshal(startAppReq)
	if err != nil {
		log.Errorf("StartApp req(%+v) resp(%+v) err(%+v)", startAppReq, reqJson, err)
		return
	}
	return ApiRequest(c, key, secret, host, _startPath, string(reqJson))
}

// AppHeart app心跳
func AppHeart(c context.Context, key, secret, host, gameId string) (resp BaseResp, err error) {
	appHeartbeatReq := AppHeartbeatReq{
		GameId: gameId,
	}
	reqJson, err := json.Marshal(appHeartbeatReq)
	if err != nil {
		log.Errorf("AppHeart req(%+v) resp(%+v) err(%+v)", appHeartbeatReq, reqJson, err)
		return
	}
	return ApiRequest(c, key, secret, host, _heartPath, string(reqJson))
}

// EndApp 关闭app
func EndApp(c context.Context, key, secret, host, gameId string, appId int64) (resp BaseResp, err error) {
	endAppReq := EndAppRequest{
		GameId: gameId,
		AppId:  appId,
	}
	reqJson, err := json.Marshal(endAppReq)
	if err != nil {
		log.Errorf("EndApp req(%+v) resp(%+v) err(%+v)", endAppReq, reqJson, err)
		return
	}
	return ApiRequest(c, key, secret, host, _endPath, string(reqJson))
}
