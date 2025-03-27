package websocket

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

// StartWebsocket 启动长连
func (w *Websocket) StartWebsocket(c context.Context, gCtx context.Context, wsAddr, authBody string, doMsg protoLogic) (err error) {
	w.log.WithContext(c).Info("StartWebsocket wsAddr:", wsAddr)
	// 建立连接
	conn, _, err := websocket.DefaultDialer.Dial(wsAddr, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.ws = &WebsocketClient{
		conn:      conn,
		msgBuf:    make(chan *Proto, 1024),
		dispather: make(map[int32]protoLogic),
	}

	// 注册分发处理函数
	w.ws.dispather[OP_AUTH_REPLY] = w.ws.authResp
	w.ws.dispather[OP_HEARTBEAT_REPLY] = w.ws.heartBeatResp
	w.ws.dispather[OP_SEND_SMS_REPLY] = w.ws.msgResp
	if doMsg != nil {
		w.ws.dispather[OP_SEND_SMS_REPLY] = doMsg
	}

	// 发送鉴权信息
	err = w.ws.sendAuth(authBody)
	if err != nil {
		w.log.WithContext(c).Error("StartWebsocket sendAuth err:", err)
		return
	}

	// todo: 结束后需要主动回收goroutine嘛？
	// 读取信息
	go w.ws.ReadMsg(gCtx)

	// 处理信息
	go w.ws.DoEvent(gCtx)

	return
}

// ReadMsg 读取长连信息
func (wc *WebsocketClient) ReadMsg(c context.Context) {
	for {
		select {
		case <-c.Done():
			return
		default:
			retProto := &Proto{}
			_, buf, err := wc.conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
				continue
			}
			retProto.PacketLength = int32(binary.BigEndian.Uint32(buf[PackOffset:HeaderOffset]))
			retProto.HeaderLength = int16(binary.BigEndian.Uint16(buf[HeaderOffset:VerOffset]))
			retProto.Version = int16(binary.BigEndian.Uint16(buf[VerOffset:OperationOffset]))
			retProto.Operation = int32(binary.BigEndian.Uint32(buf[OperationOffset:SeqIdOffset]))
			retProto.SequenceId = int32(binary.BigEndian.Uint32(buf[SeqIdOffset:]))
			if retProto.PacketLength < 0 || retProto.PacketLength > MaxPackSize {
				continue
			}
			if retProto.HeaderLength != RawHeaderSize {
				continue
			}
			if bodyLen := int(retProto.PacketLength - int32(retProto.HeaderLength)); bodyLen > 0 {
				retProto.Body = buf[retProto.HeaderLength:retProto.PacketLength]
			} else {
				continue
			}
			retProto.BodyMuti = [][]byte{retProto.Body}
			if len(retProto.BodyMuti) > 0 {
				retProto.Body = retProto.BodyMuti[0]
			}
			wc.msgBuf <- retProto
		}
	}
}

// DoEvent 处理信息
func (wc *WebsocketClient) DoEvent(c context.Context) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-c.Done():
			return
		case p := <-wc.msgBuf:
			if p == nil {
				continue
			}
			if wc.dispather[p.Operation] == nil {
				continue
			}
			err := wc.dispather[p.Operation](p)
			if err != nil {
				continue
			}
		case <-ticker.C:
			wc.sendHeartBeat()
		}
	}
}

// sendAuth 发送鉴权
func (wc *WebsocketClient) sendAuth(authBody string) (err error) {
	p := &Proto{
		Operation: OP_AUTH,
		Body:      []byte(authBody),
	}
	return wc.sendMsg(p)
}

// sendHeartBeat 发送心跳
func (wc *WebsocketClient) sendHeartBeat() {
	if !wc.authed {
		return
	}
	msg := &Proto{}
	msg.Operation = OP_HEARTBEAT
	msg.SequenceId = wc.sequenceId
	wc.sequenceId++
	err := wc.sendMsg(msg)
	if err != nil {
		return
	}
	log.Println("[WebsocketClient | sendHeartBeat] seq:", msg.SequenceId)
}

// sendMsg 发送信息
func (wc *WebsocketClient) sendMsg(msg *Proto) (err error) {
	dataBuff := &bytes.Buffer{}
	packLen := int32(RawHeaderSize + len(msg.Body))
	msg.HeaderLength = RawHeaderSize
	binary.Write(dataBuff, binary.BigEndian, packLen)
	binary.Write(dataBuff, binary.BigEndian, int16(RawHeaderSize))
	binary.Write(dataBuff, binary.BigEndian, msg.Version)
	binary.Write(dataBuff, binary.BigEndian, msg.Operation)
	binary.Write(dataBuff, binary.BigEndian, msg.SequenceId)
	binary.Write(dataBuff, binary.BigEndian, msg.Body)
	err = wc.conn.WriteMessage(websocket.BinaryMessage, dataBuff.Bytes())
	if err != nil {
		log.Println("[WebsocketClient | sendMsg] send msg err:", msg)
		return
	}
	return
}

// authResp 鉴权处理函数
func (wc *WebsocketClient) authResp(msg *Proto) (err error) {
	resp := &AuthRespParam{}
	err = json.Unmarshal(msg.Body, resp)
	if err != nil {
		return
	}
	if resp.Code != 0 {
		return
	}
	wc.authed = true
	log.Println("[WebsocketClient | authResp] auth success")
	return
}

// heartBeatResp  心跳结果
func (wc *WebsocketClient) heartBeatResp(msg *Proto) (err error) {
	log.Println("[WebsocketClient | heartBeatResp] get HeartBeat resp", msg.Body)
	return
}

// msgResp 可以这里做回调
func (wc *WebsocketClient) msgResp(msg *Proto) (err error) {
	// https://open-live.bilibili.com/document/f9ce25be-312e-1f4a-85fd-fef21f1637f8#h1-u76F4u64ADu95F4u6570u636E
	for index, cmd := range msg.BodyMuti {
		log.Printf("[WebsocketClient | msgResp] recv MsgResp index:%d ver:%d cmd:%s", index, msg.Version, string(cmd))

	}
	return
}
