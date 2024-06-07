package action

import (
	"encoding/json"
	"fmt"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

func UserGetUsers(options jsonrpc.Options) (json.RawMessage, error) {
	var users []entity.User
	options.Conn.Find(&users)

	fmt.Println("USERS:", users)
	result, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return result, nil
}
