package service

import (
	"time"

	"database/sql"

	pg "../postgres"
	vk "../vkapi"
)

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
