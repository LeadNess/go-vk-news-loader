package service

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	pg "github.com/vnkrtv/go-vk-news-loader/pkg/postgres"
	vk "github.com/vnkrtv/go-vk-news-loader/pkg/vkapi"
)

type Config struct {
	PGUser   string `json:"pguser"`
	PGPass   string `json:"pgpass"`
	PGName   string `json:"pgname"`
	PGHost   string `json:"pghost"`
	PGPort   string `json:"pgport"`
	VKToken  string `json:"vktoken"`
	Interval int    `json:"interval"`
}

func ParseVKGroup(vkGroup vk.VKGroup) pg.Group {
	return pg.Group{
		ID:           vkGroup.ID,
		ScreenName:   vkGroup.ScreenName,
		Name:         vkGroup.Name,
		MembersCount: vkGroup.MembersCount,
	}
}

func ParseVKWall(vkWall vk.VKWall, groupScreenName string) []pg.Post {
	var posts []pg.Post
	for _, post := range vkWall.Items {
		if len(post.Attachments) != 0 &&
			post.Attachments[0].Link.Title != "" &&
			post.Attachments[0].Link.Description != "" {
			group := sql.NullString{
				String: groupScreenName,
				Valid:  true,
			}
			post := pg.Post{
				ID:              post.ID,
				GroupScreenName: group,
				Date:            time.Unix(int64(post.Date), 0),
				Title:           post.Attachments[0].Link.Title,
				Text:            post.Attachments[0].Link.Description,
				LikesCount:      post.Likes.Count,
				ViewsCount:      post.Views.Count,
				CommentsCount:   post.Comments.Count,
				RepostsCount:    post.Reposts.Count,

			}
			posts = append(posts, post)
		}
	}
	return posts
}

func GetGroupsScreenNames(groupsPath string) ([]string, error) {
	bytes, err := ioutil.ReadFile(groupsPath)
	if err != nil {
		return nil, err
	}
	var groupsScreenNames []string
	err = json.Unmarshal(bytes, &groupsScreenNames)
	return groupsScreenNames, err
}

func GetConfig() (Config, error) {
	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		if i := strings.Index(env, "="); i >= 0 {
			envVars[env[:i]] = env[i+1:]
		}
	}
	interval, err := strconv.ParseInt(envVars["DATA_LOAD_INTERVAL"], 10, 64)
	if err != nil {
		return Config{}, err
	}
	config := Config{
		PGUser:   envVars["PG_USER"],
		PGPass:   envVars["PG_PASS"],
		PGName:   envVars["PG_NAME"],
		PGHost:   envVars["PG_HOST"],
		PGPort:   envVars["PG_PORT"],
		VKToken:  envVars["VK_TOKEN"],
		Interval: int(interval),
	}
	return config, err
}
