package action

import (
	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
	"encoding/json"
	"fmt"
)

type GroupGetGroupsDataGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UsersCount  int    `json:"users"`
	LinksCount  int    `json:"links"`
}

type GroupGetGroupsData struct {
	My []GroupGetGroupsDataGroup `json:"my"`
}

type GetGroupsResponse struct {
	My  []entity.Group `json:"my"`
	All []entity.Group `json:"all"`
}

func GroupGetGroups(options jsonrpc.Options) (json.RawMessage, error) {
	my := make([]entity.Group, 0)

	options.Conn.Model(&entity.User{ID: options.UserID}).Order("id DESC").Association("Groups").Find(&my)

	ggr := &GetGroupsResponse{My: my}
	fmt.Println(ggr)
	return json.Marshal(ggr)
}
