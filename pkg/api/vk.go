package api

import (
	"github.com/go-vk-api/vk"
)

type VKAPi struct {
	api *vk.Client
}

func NewVKApi(token string) (*VKAPi, error) {
	api, err := vk.NewClientWithOptions(
		vk.WithToken(token),
	)
	return &VKAPi{
		api: api,
	}, err
}

func (a *VKAPi) WallGet(domain string) (VKWall, error) {
	var wall VKWall
	err := a.api.CallMethod("wall.get", vk.RequestParams{
		"domain": domain,
		"count": 100,
	}, &wall)
	return wall, err
}
