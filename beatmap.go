package osu_go_client

import (
	"fmt"
	"github.com/pkg/errors"
)

type BeatmapAPI struct {
	*OsuAPI
}

// Get general beatmap information by beatmapId
func (b *BeatmapAPI) Get(id uint) (*SingleBeatmap, error) {
	json := SingleBeatmap{}

	req := b.client.
		Request().
		Path(fmt.Sprint("/api/v2/beatmap/", id)).
		Use(b.bearerMiddleware())

	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	if !res.Ok {
		return nil, errors.New(res.RawResponse.Status)
	}

	if err := res.JSON(&json); err != nil {
		return nil, err
	}

	return &json, nil
}
