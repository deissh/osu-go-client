package main

import (
	"github.com/deissh/osu-go-client"
	"log"
	"os"
	"sync"
)

func fetchSet(wg *sync.WaitGroup, client osu_go_client.OsuAPI, beatMapId uint) {
	defer wg.Done()

	data, err := client.BeatmapSet.Get(beatMapId)
	if err != nil {
		log.Print(err)
	} else {
		log.Println(data.ID)
		log.Println(data.Title)
	}
}

func main() {
	api := osu_go_client.WithAccessToken(
		os.Getenv("access_token"),
		os.Getenv("refresh_token"),
	)

	data, err := api.BeatmapSet.Get(23416)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(data.ID)
	log.Println(data.Title)

	var wg sync.WaitGroup
	beatMapIds := []uint{816264, 765778, 23416, 1614054}
	for _, id := range beatMapIds {
		wg.Add(1)

		go fetchSet(&wg, api, id)
	}
	wg.Wait()
}
