# osu!lazer api client

## Features

1. Lazer version
2. Thread-safe
3. Auto refresh access_token
4. ...

## Using

```go
import client "github.com/deissh/osu-go-client"
```

```go
// create new client with username & password
// api := client.WithBasicAuth(
//	   os.Getenv("username"),
//     os.Getenv("password"),
// )

// or with access_token and refresh_token
api := client.WithAccessToken(
    os.Getenv("access_token"),
    os.Getenv("refresh_token"),
)

data, err := api.BeatmapSet.Get(23416)
if err != nil {
    // error
    return
}

// data.ID
// data.Title
// ...
```

You can also use one client in different goroutines

```go
api := osu_go_client.WithAccessToken(
    os.Getenv("access_token"),
    os.Getenv("refresh_token"),
)

beatMapIds := []uint{141515, 514551, 23416, 261441}

// run new goroutines for each beatMapId
for _, id := range beatMapIds {
    go func(beatMapId uint) {
        data, err := api.BeatmapSet.Get(23416)
        if err != nil {
            log.Fatal(err)
        }

        log.Println(data)
    }(id)
}
```
