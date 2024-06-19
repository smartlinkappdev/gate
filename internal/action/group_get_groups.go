package action

import (
	"encoding/json"
	"fmt"

	"cmd/gate/main.go/internal/entity"
	"cmd/gate/main.go/internal/jsonrpc"
)

type GetGroupsResponse struct {
	My  []entity.Group `json:"my"`
	All []entity.Group `json:"all"`
}

func GroupGetGroups(options jsonrpc.Options) (json.RawMessage, error) {
	var groups []entity.Group
	options.Conn.Model(&entity.Group{}).Order("id DESC").Preload("Users").Find(&groups)

	my := make([]entity.Group, 0)
	all := make([]entity.Group, 0)

	options.Conn.Model(&entity.User{ID: options.UserID}).Association("Groups").Find(&my)

	//options.Conn.Where("owner_id = ?", options.UserID).Order("id DESC").Find(&my)
	options.Conn.Model(&entity.Group{}).Order("id DESC").Find(&all)

	ggr := &GetGroupsResponse{My: my, All: all}
	fmt.Println(ggr)
	return json.Marshal(ggr)
}
