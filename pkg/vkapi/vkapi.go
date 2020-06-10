package vkapi

import (
	"fmt"
	"github.com/go-vk-api/vk"
	"strconv"
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

func (a *VKAPi) GetGroupsPosts(domains []string, postsCount int) (map[string]VKWall, error) {
	groupsDomains := make([]string, len(domains))
	for i, str := range domains {
		groupsDomains[i] = fmt.Sprintf("%s", strconv.Quote(str))
		if i != len(domains) - 1{
			groupsDomains[i] += ","
		}
	}
	var response []VKWall
	code := `
        var domains = %s;
		var res = [];
		var i = 0;
		while (i < domains.length) {
			var posts = API.wall.get({domain: domains[i], count: %d});
			res.push(posts);
			i = i + 1; 
		}
		return res;`
	err := a.api.CallMethod("execute", vk.RequestParams{
		"code": fmt.Sprintf(code, groupsDomains, postsCount),
	}, &response)
	wallMap := make(map[string]VKWall, len(domains))
	for i, wall := range response {
		wallMap[domains[i]] = wall
	}
	return wallMap, err
}