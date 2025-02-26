package live

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"xiuyiPro/internal/conf"
	"xiuyiPro/internal/domain/websocket"
)

type Live struct {
	cfg   *conf.Application_Live
	log   *log.Helper
	WsDao *websocket.Websocket
}

func New(c *conf.Application_Live, logger *log.Helper) (*Live, error) {
	if c == nil {
		return nil, errors.New("config for Live is nil")
	}
	l := &Live{
		cfg: c,
		log: logger,
	}
	var err error
	if l.WsDao, err = websocket.New(l.log); err != nil {
		return nil, err
	}
	return l, nil
}
