package osu_go_client

import (
	"github.com/bxcodec/faker/v3"
	"github.com/deissh/go-utils"
	"github.com/dgrijalva/jwt-go"
	mock "gopkg.in/h2non/gentleman-mock.v2"
	"reflect"
	"testing"
	"time"
)

func genFakeBeatmap(t *testing.T) SingleBeatmap {
	b := SingleBeatmap{}

	err := faker.FakeData(&b)
	if err != nil {
		t.Error(err)
	}

	return b
}

func TestBeatmapAPI_Get(t *testing.T) {
	defer mock.Disable()

	type args struct {
		id uint
	}
	type response struct {
		status int
		json   interface{}
	}

	tests := []struct {
		name     string
		args     args
		response response
		wantErr  bool
	}{
		{
			name: "get with id",
			args: args{id: 1},
			response: response{
				status: 200,
				json:   genFakeBeatmap(t),
			},
			wantErr: false,
		},
		{
			name: "get with invalid id",
			args: args{id: 1},
			response: response{
				status: 404,
				json:   nil,
			},
			wantErr: true,
		},
		{
			name: "get with invalid auth",
			args: args{id: 1},
			response: response{
				status: 401,
				json:   nil,
			},
			wantErr: true,
		},
		{
			name: "get with invalid body",
			args: args{id: 1},
			response: response{
				status: 200,
				json:   "invalid json response",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new token object, specifying signing method and the claims
			// you would like it to contain.
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().UTC().Add(time.Hour * 24).Unix(),
			})

			// Sign and get the complete encoded token as a string using the secret
			accessToken, err := token.SignedString([]byte("some_super_ultra_secret"))
			if err != nil {
				t.Error(err)
			}

			// creating test client
			testClient := WithAccessToken(accessToken, "somerefreshtoken")
			// Register the mock plugin at client level
			testClient.client.Use(mock.Plugin)

			// Configure the mock via gock
			mock.New(APIDomain).
				Get("/api/v2/beatmap/*").
				Reply(tt.response.status).
				JSON(tt.response.json)

			got, err := testClient.Beatmap.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(utils.ToStringMap(got), utils.ToStringMap(tt.response.json)) {
				t.Errorf("Get() got = %v, want %v", got, &tt.response.json)
			}
		})
	}
}
