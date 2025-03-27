package errorcode

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ParamInvalid     = errors.New(1219, "", "参数错误")
	ServiceError     = errors.New(1220, "", "服务异常")
	GameNotFound     = errors.New(1221, "", "未找到游戏")
	NoAvailableIdiom = errors.New(1222, "", "无可用题目")
	RankingNotFound  = errors.New(1223, "", "未找到排行榜")
	StartAppFailed   = errors.New(1224, "", "启动应用失败")
	CloseAppFailed   = errors.New(1225, "", "关闭应用失败")
	CreateWsFailed   = errors.New(1226, "", "创建长连接失败")
)
