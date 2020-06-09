package main

import (
	"../pkg/api"
	"fmt"
	"log"
	"time"
	"os"
)

func main() {
	token := os.Getenv("VK_TOKEN")
	session, err := api.NewVKApi(token)
	if err != nil {
		log.Fatal(err)
	}
	groupsDomains := []string{"meduzaproject", "ria", "kommersant_ru", "tj", "rbc"}
	groupsWall, err := session.GetGroupsPosts(groupsDomains, 3)

	if err != nil {
		log.Fatal(err)
	} else {
		for domain, wall := range groupsWall {
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
}