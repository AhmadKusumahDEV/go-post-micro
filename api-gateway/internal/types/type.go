package types

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound      = errors.New("not found Data")
	ErrConnetion     = errors.New("connection error")
	ErrDecode        = errors.New("decode error")
	ErrEncode        = errors.New("encode error")
	ErrRead          = errors.New("error read data to byte, %w")
	ErrCreateRequest = errors.New("create request error, %w")
	ErrSendRequest   = errors.New("send request error, %w")
	ErrConvertData   = errors.New("convert data error, %w")
	ErrRedis         = errors.New("redis error, %w")
	ErrRepository    = errors.New("repository Error, %w")
)

type Midel struct {
	Handler http.Handler
}

func (t *Midel) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

type Header string
type ChannelGroup map[string]chan bool

const (
	HeaderKey     Header = "header"
	RespKeyHeader Header = "resp"
	PostKeyCtx    Header = "post"
)
