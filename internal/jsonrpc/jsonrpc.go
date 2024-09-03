package jsonrpc

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"log/slog"

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
	Log     *slog.Logger
	Conn    *gorm.DB
	Conn2   *gorm.DB
	Conn3ch *gorm.DB
	UserID  int
	Params  json.RawMessage
	Cash    map[string]json.RawMessage
	Rdb     *redis.Client
}

type Method func(options Options) (json.RawMessage, error)
