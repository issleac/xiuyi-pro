package live

import (
	"encoding/json"
	"xiuyiPro/internal/domain/websocket"
)

const (
	AcceptHeader              = "Accept"
	ContentTypeHeader         = "Content-Type"
	AuthorizationHeader       = "Authorization"
	JsonType                  = "application/json"
	BiliVersion               = "1.0"
	HmacSha256                = "HMAC-SHA256"
	BiliTimestampHeader       = "x-bili-timestamp"
	BiliSignatureMethodHeader = "x-bili-signature-method"
	BiliSignatureNonceHeader  = "x-bili-signature-nonce"
	BiliAccessKeyIdHeader     = "x-bili-accesskeyid"
	BiliSignVersionHeader     = "x-bili-signature-version"
	BiliContentMD5Header      = "x-bili-content-md5"

	OpenLiveCmdDm    = "LIVE_OPEN_PLATFORM_DM"
	OpenLiveCmdGift  = "LIVE_OPEN_PLATFORM_SEND_GIFT"
	OpenLiveCmdClose = "LIVE_OPEN_PLATFORM_INTERACTION_END"
)

type CommonHeader struct {
	ContentType       string
	ContentAcceptType string
	Timestamp         string
	SignatureMethod   string
	SignatureVersion  string
	Authorization     string
	Nonce             string
	AccessKeyId       string
	ContentMD5        string
}

type BaseResp struct {
	Code      int64           `json:"code"`
	Message   string          `json:"message"`
	RequestId string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
}

type StartAppRequest struct {
	// 主播身份码
	Code string `json:"code"`
	// 项目id
	AppId int64 `json:"app_id"`
}

type StartAppRespData struct {
	// 场次信息
	GameInfo GameInfo `json:"game_info"`
	// 长连信息
	WebsocketInfo WebSocketInfo `json:"websocket_info"`
	// 主播信息
	AnchorInfo AnchorInfo `json:"anchor_info"`
}

type GameInfo struct {
	GameId string `json:"game_id"`
}

type WebSocketInfo struct {
	//  长连使用的请求json体 第三方无需关注内容,建立长连时使用即可
	AuthBody string `json:"auth_body"`
	//  wss 长连地址
	WssLink []string `json:"wss_link"`
}

type AnchorInfo struct {
	//主播房间号
	RoomId int64 `json:"room_id"`
	//主播昵称
	Uname string `json:"uname"`
	//主播头像
	Uface string `json:"uface"`
	//主播uid
	Uid int64 `json:"uid"`
	//主播open_id
	OpenId string `json:"open_id"`
}

type EndAppRequest struct {
	// 场次id
	GameId string `json:"game_id"`
	// 项目id
	AppId int64 `json:"app_id"`
}

type AppHeartbeatReq struct {
	// 主播身份码
	GameId string `json:"game_id"`
}

type RoomGame struct {
	RoomId   string
	UpCodeId string
	Ws       *websocket.Websocket
}

type OpenPlatformDM struct {
	RoomId                 int64  `json:"room_id"`
	Uid                    int64  `json:"uid"`
	OpenId                 string `json:"open_id"`
	Uname                  string `json:"uname"`
	Msg                    string `json:"msg"`
	MsgId                  string `json:"msg_id"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
	GuardLevel             int    `json:"guard_level"`
	Timestamp              int64  `json:"timestamp"`
	Uface                  string `json:"uface"`
	EmojiImgUrl            string `json:"emoji_img_url"`
	DmType                 int    `json:"dm_type"`
	GloryLevel             int    `json:"glory_level"`
	ReplyOpenId            string `json:"reply_open_id"`
	ReplyUname             string `json:"reply_uname"`
	IsAdmin                int    `json:"is_admin"`
}

type OpenPlatformGift struct {
	RoomId                 int64  `json:"room_id"`
	Uid                    int64  `json:"uid"`
	OpenId                 string `json:"open_id"`
	Uname                  string `json:"uname"`
	Uface                  string `json:"uface"`
	GiftId                 int    `json:"gift_id"`
	GiftName               string `json:"gift_name"`
	GiftNum                int    `json:"gift_num"`
	Price                  int    `json:"price"`
	Paid                   bool   `json:"paid"`
	FansMedalLevel         int    `json:"fans_medal_level"`
	FansMedalName          string `json:"fans_medal_name"`
	FansMedalWearingStatus bool   `json:"fans_medal_wearing_status"`
	GuardLevel             int    `json:"guard_level"`
	Timestamp              int    `json:"timestamp"`
	MsgId                  string `json:"msg_id"`
	AnchorInfo             struct {
		Uid    int    `json:"uid"`
		OpenId string `json:"open_id"`
		Uname  string `json:"uname"`
		Uface  string `json:"uface"`
	} `json:"anchor_info"`
	GiftIcon  string `json:"gift_icon"`
	ComboGift bool   `json:"combo_gift"`
	ComboInfo struct {
		ComboBaseNum int    `json:"combo_base_num"`
		ComboCount   int    `json:"combo_count"`
		ComboId      string `json:"combo_id"`
		ComboTimeout int    `json:"combo_timeout"`
	} `json:"combo_info"`
}
