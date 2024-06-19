package action

import (
	"encoding/json"
	"fmt"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

func UserGetUser(options jsonrpc.Options) (json.RawMessage, error) {
	user := entity.User{}
	options.Conn.Model(&user).Where("id = ?", options.UserID).First(&user)

	fmt.Println("USER:", user)
	result, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
