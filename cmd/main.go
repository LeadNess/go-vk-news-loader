package main

import (
	"log"
	"os"

	"../pkg/service"
)

var (
	vkToken    string
	pgUser     string
	pgPassword string
	pgHost     string
	pgPort     string
	pgDBName   string
)

func init() {
	vkToken = os.Getenv("VK_TOKEN")
	pgUser     = os.Getenv("PG_USER")
	pgPassword = os.Getenv("PG_PASSWORD")
	pgHost     = os.Getenv("PG_HOST")
	pgPort     = os.Getenv("PG_PORT")
	pgDBName   = os.Getenv("PG_DBNAME")
}

func main() {
	service, err := service.NewNewsService(
		vkToken, pgUser, pgPassword, pgHost, pgPort, pgDBName)
	if err != nil {
		log.Fatal(err)
	}
	groupsScreenNames := []string{"meduzaproject", "ria", "kommersant_ru", "tj", "rbc"}
	if err := service.InitDB(); err != nil {
		log.Fatal(err)
	}
	if err := service.AddNewsSources(groupsScreenNames); err != nil {
		log.Fatal(err)
	}
	if err := service.LoadNews(groupsScreenNames, 100); err != nil {
		log.Fatal(err)
	}
}