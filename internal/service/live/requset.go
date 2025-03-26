package live

import (
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/monaco-io/request"
	"github.com/pkg/errors"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ApiRequest http request demo方法
func ApiRequest(c context.Context, accessKey, accessSecret, host, requestUrl, reqJson string) (resp BaseResp, err error) {
	resp = BaseResp{}
	header := &CommonHeader{
		ContentType:       JsonType,
		ContentAcceptType: JsonType,
		Timestamp:         strconv.FormatInt(time.Now().Unix(), 10),
		SignatureMethod:   HmacSha256,
		SignatureVersion:  BiliVersion,
		Authorization:     "",
		Nonce:             strconv.FormatInt(time.Now().UnixNano(), 10), //用于幂等,记得替换
		AccessKeyId:       accessKey,
		ContentMD5:        Md5(reqJson),
	}
	header.Authorization = CreateSignature(header, accessSecret)
	cli := request.Client{
		Method: "POST",
		URL:    fmt.Sprintf("%s%s", host, requestUrl),
		Header: header.ToMap(),
		String: reqJson,
	}
	cliResp := cli.Send().Scan(&resp)
	if !cliResp.OK() {
		err = errors.Wrapf(cliResp.Error(), "[error] req:%+v resp:%+v", reqJson, resp)
		log.Errorf("ApiRequest cli(%+v) cliResp(%+v) err(%+v)", cli, cliResp, err)
		return
	}
	return
}

// CreateSignature 生成Authorization加密串
func CreateSignature(header *CommonHeader, accessKeySecret string) string {
	sStr := header.ToSortedString()
	return HmacSHA256(accessKeySecret, sStr)
}

// Md5 md5加密
func Md5(str string) (md5str string) {
	data := []byte(str)
	has := md5.Sum(data)
	md5str = fmt.Sprintf("%x", has)
	return md5str
}

// HmacSHA256 HMAC-SHA256算法
func HmacSHA256(key string, data string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// ToMap 所有字段转map<string, string>
func (h *CommonHeader) ToMap() map[string]string {
	return map[string]string{
		BiliTimestampHeader:       h.Timestamp,
		BiliSignatureMethodHeader: h.SignatureMethod,
		BiliSignatureNonceHeader:  h.Nonce,
		BiliAccessKeyIdHeader:     h.AccessKeyId,
		BiliSignVersionHeader:     h.SignatureVersion,
		BiliContentMD5Header:      h.ContentMD5,
		AuthorizationHeader:       h.Authorization,
		ContentTypeHeader:         h.ContentType,
		AcceptHeader:              h.ContentAcceptType,
	}
}

// ToSortMap 参与加密的字段转map<string, string>
func (h *CommonHeader) ToSortMap() map[string]string {
	return map[string]string{
		BiliTimestampHeader:       h.Timestamp,
		BiliSignatureMethodHeader: h.SignatureMethod,
		BiliSignatureNonceHeader:  h.Nonce,
		BiliAccessKeyIdHeader:     h.AccessKeyId,
		BiliSignVersionHeader:     h.SignatureVersion,
		BiliContentMD5Header:      h.ContentMD5,
	}
}

// ToSortedString 生成需要加密的文本
func (h *CommonHeader) ToSortedString() (sign string) {
	hMap := h.ToSortMap()
	var hSil []string
	for k := range hMap {
		hSil = append(hSil, k)
	}
	sort.Strings(hSil)
	for _, v := range hSil {
		sign += v + ":" + hMap[v] + "\n"
	}
	sign = strings.TrimRight(sign, "\n")
	return
}
