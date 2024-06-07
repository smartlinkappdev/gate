package jsonrpc

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Request struct {
	Token  string          `json:"token"`
	Method string          `json:"method"`
	Params json.RawMessage `json:"params"`
}

type Response struct {
	Auth   bool            `json:"auth"`
	Error  string          `json:"error"`
	Result json.RawMessage `json:"result"`
}

type Options struct {
	Conn   *gorm.DB
	UserID int
	Params json.RawMessage
}

type Method func(options Options) (json.RawMessage, error)
