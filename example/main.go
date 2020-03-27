package main

import (
	"github.com/deissh/osu-go-client"
	"log"
	"os"
)

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

	beatMapIds := []uint{141515, 514551, 23416, 261441}
	for _, id := range beatMapIds {
		go func(beatMapId uint) {
			data, err := api.BeatmapSet.Get(23416)
			if err != nil {
				log.Fatal(err)
			}

			log.Println(data)
		}(id)
	}

	log.Println(data.ID)
	log.Println(data.Title)
}
