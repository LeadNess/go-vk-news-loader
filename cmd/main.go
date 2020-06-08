package main

import (
	"fmt"
	"../pkg/api"
	"log"
	"os"
)

func main() {
	token := os.Getenv("VK_TOKEN")
	session, err := api.NewVKApi(token)
	if err != nil {
		log.Fatal(err)
	}
	wall, err := session.WallGet("meduzaproject")
	if err != nil {
		log.Fatal(err)
	} else {
		for _, post := range wall.Items {
			fmt.Printf("%s\n%s\n\n", post.Attachments[0].Link.Title, post.Attachments[0].Link.Description)
		}
	}
}