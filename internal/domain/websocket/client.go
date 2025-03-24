package websocket

import (
	"github.com/go-kratos/kratos/v2/log"
)

type Websocket struct {
	log *log.Helper
	ws  *WebsocketClient
}

func New(logger *log.Helper) (*Websocket, error) {
	return &Websocket{
		log: logger,
	}, nil
}
