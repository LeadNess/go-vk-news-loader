package service

import (
	"fmt"
	"time"

	pg "../postgres"
	vk "../vkapi"
)

type NewsLoader interface {
	LoadNews(groupsDomains []string, count int) error
}

type NewsService struct {
	db          *pg.Storage
	api         *vk.VKAPi
	lastUpdate  time.Time
	loadedPosts []vk.VKPost
}

func NewNewsLoaderService(vkToken, pgUser, pgPass, pgHost, pgPort, pgDBName string) (*NewsService, error) {
	db, err := pg.OpenConnection(pgUser, pgPass, pgHost, pgPort, pgDBName)
	if err != nil {
		return nil, err
	}
	api, err := vk.NewVKApi(vkToken)
	if err != nil {
		return nil, err
	}
	return &NewsService{
		db:  db,
		api: api,
	}, err
}

func (s *NewsService) AddNewsSource(groupScreenName string) error {
	s.db.InsertGroup()
}

func (s *NewsService) LoadNews(groupsDomains []string, count int) error {
	mapNews, err := s.api.GetGroupsPosts(groupsDomains, count)
	if err != nil {
		return err
	}
	for domain, wall := range mapNews {
		fmt.Printf("Group: %s\n\n", domain)
		for _, post := range wall.Items {
			if len(post.Attachments) != 0 &&
				post.Attachments[0].Link.Title != "" &&
				post.Attachments[0].Link.Description != "" {
				fmt.Printf("Title: %s\n\nDescription: %s\n\nDate: %s\nLikes: %d\nViews: %d\nComments: %d\n\n\n",
					post.Attachments[0].Link.Title, post.Attachments[0].Link.Description,
					time.Unix(post.Date, 0), post.Likes.Count, post.Views.Count, post.Comments.Count)
			}
		}
	}
}