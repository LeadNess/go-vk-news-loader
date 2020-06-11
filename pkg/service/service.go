package service

import (
	"time"

	"github.com/pkg/errors"

	pg "../postgres"
	vk "../vkapi"
)

var (
	IncorrectScreenName = errors.New("incorrect group screen name")
)

type NewsLoader interface {
	LoadNews(groupsDomains []string, count int) error
}

type NewsService struct {
	db          *pg.Storage
	vkApi       *vk.VKAPi
	lastUpdate  time.Time
	latestPosts map[string]pg.Post
}

func NewNewsService(vkToken, pgUser, pgPass, pgHost, pgPort, pgDBName string) (*NewsService, error) {
	db, err := pg.OpenConnection(pgUser, pgPass, pgHost, pgPort, pgDBName)
	if err != nil {
		return nil, err
	}
	api, err := vk.NewVKApi(vkToken)
	if err != nil {
		return nil, err
	}
	return &NewsService{
		db:          db,
		vkApi:       api,
		latestPosts: make(map[string]pg.Post),
	}, err
}

func (s *NewsService) InitDB() error {
	return s.db.CreateSchema()
}

func (s *NewsService) AddNewsSource(groupScreenName string) error {
	vkGroups, err := s.vkApi.GetGroups([]string{groupScreenName})
	if err != nil {
		return err
	} else if len(vkGroups) == 0 {
		return IncorrectScreenName
	}
	group := ParseVKGroup(vkGroups[0])
	return s.db.InsertGroup(group)
}

func (s *NewsService) AddNewsSources(groupsScreenNames []string) error {
	vkGroups, err := s.vkApi.GetGroups(groupsScreenNames)
	if err != nil {
		return err
	} else if len(vkGroups) == 0 {
		return IncorrectScreenName
	}
	groups := make([]pg.Group, len(vkGroups))
	for i, vkGroup := range vkGroups {
		groups[i] = ParseVKGroup(vkGroup)
	}
	return s.db.InsertGroups(groups)
}

func (s *NewsService) LoadNews(groupsDomains []string, count int) error {
	mapNews, err := s.vkApi.GetGroupsPosts(groupsDomains, count)
	if err != nil {
		return err
	}
	for group, wall := range mapNews {
		posts := ParseVKWall(wall, group)
		if _, ok := s.latestPosts[group]; ok {
			latestPost := s.latestPosts[group]
			for i, post := range posts {
				if post.ID == latestPost.ID {
					posts = posts[:i]
					break
				}
			}
		}
		s.latestPosts[group] = posts[0]
		if err := s.db.InsertPosts(posts); err != nil {
			return err
		}
	}
	return err
}
