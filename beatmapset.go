package osu_go_client

import (
	"errors"
	"fmt"
)

type BeatmapSetAPI struct {
	*OsuAPI
}

func (b *BeatmapSetAPI) Get(id uint) (*BeatmapSetFull, error) {
	json := BeatmapSetFull{}

	req := b.client.
		Request().
		Path(fmt.Sprint("/api/v2/beatmapsets/", id)).
		Use(b.bearerMiddleware())

	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New("request status: " + res.RawResponse.Status)
	}

	if err := res.JSON(&json); err != nil {
		return nil, err
	}

	return &json, nil
}
