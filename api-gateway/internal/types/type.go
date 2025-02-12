package types

import (
	"errors"
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

type Header string

const (
	HeaderKey  Header = "header"
	RespHeader Header = "resp"
)
