package main

import (
	"../pkg/vkapi"
	"../pkg/postgres"
	"flag"
	"fmt"
	"log"
	"time"
	"os"
)

var (
	user = flag.String("u", "postgres", "Postgres user")
	password = flag.String("p", "password", "Postgres user password")
	host = flag.String("h", "172.17.0.2", "Postgres host")
	port = flag.String("P", "5432", "Postgres port")
	dbName = flag.String("n", "vknews", "Postgres DB name")
)

func main() {
	token := os.Getenv("VK_TOKEN")
	session, err := vkapi.NewVKApi(token)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := postgres.OpenConnection(*user, *password, *host, *port, *dbName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", conn.Conn)

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