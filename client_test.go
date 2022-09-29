//go:build test
// +build test

package osu_go_client

import (
	"log"
	"os"
)

func ExampleWithAccessToken() {
	api := WithAccessToken(
		os.Getenv("access_token"),
		os.Getenv("refresh_token"),
	)

	data, err := api.BeatmapSet.Get(23416)
	if err != nil {
		log.Fatal(err)
		return
	}
	// Output: 765778
	log.Println(data.ID)
	// Output: Make a Move (Speed Up Ver.)
	log.Println(data.Title)
}
