package main

import (
	"../pkg/service"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	vkToken      = os.Getenv("VK_TOKEN")
	pgUser       = os.Getenv("PG_USER")
	pgPassword   = os.Getenv("PG_PASSWORD")
	pgHost       = os.Getenv("PG_HOST")
	pgPort       = os.Getenv("PG_PORT")
	pgDBName     = os.Getenv("PG_DBNAME")
	timeInterval = os.Getenv("TIME_INTERVAL")
	secondsCount int64
)

func main() {
	newsService, err := service.NewNewsService(
		vkToken, pgUser, pgPassword, pgHost, pgPort, pgDBName)
	if err != nil {
		log.Fatal(err)
	}
	if secondsCount, err = strconv.ParseInt(timeInterval, 10, 64); err != nil {
		log.Fatalf("error on parsing TIME_INTERVAL - %s", err)
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
	for {
		if err := newsService.LoadNews(groupsScreenNames, 10); err != nil {
			log.Println(err)
		} else {
			log.Printf("loaded some staff")
		}
		time.Sleep(time.Duration(secondsCount) * time.Second)
	}
}