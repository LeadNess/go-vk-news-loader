package main

import (
	"log"
	"os"

	"../pkg/service"
)

var (
	vkToken     = os.Getenv("VK_TOKEN")
	pgUser      = os.Getenv("PG_USER")
	pgPassword  = os.Getenv("PG_PASSWORD")
	pgHost      = os.Getenv("PG_HOST")
	pgPort      = os.Getenv("PG_PORT")
	pgDBName    = os.Getenv("PG_DBNAME")
)

func main() {
	newsService, err := service.NewNewsService(
		vkToken, pgUser, pgPassword, pgHost, pgPort, pgDBName)
	if err != nil {
		log.Fatal(err)
	}
	groupsScreenNames := []string{"meduzaproject", "ria", "kommersant_ru", "tj", "rbc"}
	if err := newsService.InitDB(); err != nil {
		log.Fatal(err)
	}
	if err := newsService.AddNewsSources(groupsScreenNames); err != nil {
		log.Fatal(err)
	}
	if err := newsService.LoadNews(groupsScreenNames, 100); err != nil {
		log.Fatal(err)
	}
}