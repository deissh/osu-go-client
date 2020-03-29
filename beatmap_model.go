package osu_go_client

import "time"

type Beatmap struct {
	ID               int64     `json:"id"`
	BeatmapsetID     int64     `json:"beatmapset_id"`
	Mode             Mode      `json:"mode"`
	ModeInt          int64     `json:"mode_int"`
	Convert          bool      `json:"convert"`
	DifficultyRating float64   `json:"difficulty_rating"`
	Version          string    `json:"version"`
	TotalLength      int64     `json:"total_length"`
	HitLength        int64     `json:"hit_length"`
	BPM              int64     `json:"bpm"`
	CS               int64     `json:"cs"`
	Drain            int64     `json:"drain"`
	Accuracy         int64     `json:"accuracy"`
	Ar               int64     `json:"ar"`
	Playcount        int64     `json:"playcount"`
	Passcount        int64     `json:"passcount"`
	CountCircles     int64     `json:"count_circles"`
	CountSliders     int64     `json:"count_sliders"`
	CountSpinners    int64     `json:"count_spinners"`
	CountTotal       int64     `json:"count_total"`
	IsScoreable      bool      `json:"is_scoreable"`
	LastUpdated      time.Time `json:"last_updated"`
	Ranked           int64     `json:"ranked"`
	Status           Status    `json:"status"`
	URL              string    `json:"url"`
	DeletedAt        time.Time `json:"deleted_at"`
	Failtimes        Failtimes `json:"failtimes"`
	MaxCombo         *int64    `json:"max_combo"`
}

type SingleBeatmap struct {
	Beatmap
	Beatmapset BeatmapSet `json:"beatmapset"`
}

type Failtimes struct {
	Fail []int64 `json:"fail"`
	Exit []int64 `json:"exit"`
}

type Mode string

const (
	Fruits Mode = "fruits"
	Mania  Mode = "mania"
	Osu    Mode = "osu"
	Taiko  Mode = "taiko"
)

type Status string

const (
	Ranked Status = "ranked"
)
