package app

import "cmd/gate/main.go/internal/config"

type InterfaceApp interface {
	Init(config *config.Config)
	Start()
}
