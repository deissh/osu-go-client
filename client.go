package osu_go_client

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/context"
	"gopkg.in/h2non/gentleman.v2/plugin"
)

var (
	// Default api domain
	// Can be change before create api client
	APIDomain = "https://osu.ppy.sh"
	// OAuth2 client id
	// Required for create access_token or refresh its
	APIClientId = "5"
	// OAuth2 client secret
	APIClientSecret = "FGc9GAtyHzeQDshWP5Ah7dega8hJACAJpQtw6OXk"
)

// Package provide a simple interface to osu!api v2.
type OsuAPI struct {
	mux sync.Mutex

	domain       string
	clientId     string
	clientSecret string

	client       *gentleman.Client
	accessToken  string
	refreshToken string

	OAuth2     OAuth2API
	Beatmap    BeatmapAPI
	BeatmapSet BeatmapSetAPI
}

// Create new api client WithAccessToken with enabled auto refreshing access_token
// by default uses APIDomain, APIClientId, APIClientSecret, variables to set the address and the client_id/client_secret
func WithAccessToken(accessToken string, refreshToken string) *OsuAPI {
	client := gentleman.New()

	api := OsuAPI{
		domain:       APIDomain,
		clientId:     APIClientId,
		clientSecret: APIClientSecret,

		client:       client,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}

	api.OAuth2 = OAuth2API{&api}
	api.Beatmap = BeatmapAPI{&api}
	api.BeatmapSet = BeatmapSetAPI{&api}

	client.BaseURL(api.domain)

	return &api
}

func WithBasicAuth(username string, password string) (*OsuAPI, error) {
	client := gentleman.New()

	api := OsuAPI{
		domain:       APIDomain,
		clientId:     APIClientId,
		clientSecret: APIClientSecret,

		client: client,
	}

	api.OAuth2 = OAuth2API{&api}
	api.Beatmap = BeatmapAPI{&api}
	api.BeatmapSet = BeatmapSetAPI{&api}

	client.BaseURL(api.domain)

	// login
	token, err := api.OAuth2.CreateToken(username, password, "*")
	if err != nil {
		return nil, errors.Wrap(err, "failed auth")
	}

	api.accessToken = token.AccessToken
	api.refreshToken = token.RefreshToken

	return &api, nil
}

// Bearer defines an authorization bearer token header in the outgoing request
func (client *OsuAPI) bearerMiddleware() plugin.Plugin {
	return plugin.NewRequestPlugin(func(ctx *context.Context, h context.Handler) {
		client.mux.Lock()

		accessToken := client.accessToken

		expiredAt, err := parseJwtExpiredAt(accessToken)
		if err != nil {
			h.Error(ctx, err)
			return
		}

		if time.Now().After(expiredAt) {
			log.Print("refreshing token")

			token, err := client.OAuth2.TokenRenew("*", client.accessToken, client.refreshToken)
			if token == nil || err != nil {
				h.Error(ctx, errors.Wrap(err, "failed token refresh"))

				client.mux.Unlock()
				return
			}

			client.accessToken = token.AccessToken
			client.refreshToken = token.RefreshToken
		}

		ctx.Request.Header.Set("Authorization", "Bearer "+client.accessToken)

		client.mux.Unlock()
		h.Next(ctx)
	})
}

// parseJwt and return claims
func parseJwtExpiredAt(token string) (time.Time, error) {
	parser := new(jwt.Parser)
	parsedToken, _, err := parser.ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return time.Now(), errors.Wrap(err, "invalid jwt format")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return time.Now(), errors.New("invalid claims format")
	}

	var tm time.Time
	switch iat := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}

	return tm, nil
}
