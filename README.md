# Gobalt

[![Go Reference](https://pkg.go.dev/badge/github.com/andresperezl/gobalt.svg)](https://pkg.go.dev/github.com/andresperezl/gobalt)

Go library to interact with [cobalt.tools API](https://cobalt.tools/)

## Usage

Get the package into your project

```sh
go get github.com/andresperezl/gobalt/v2
```
Then import it and use it in your file

```go
import (
    "github.com/andresperezl/gobalt/v2"
)


func main() {
    // Point the client to the cobalt instance you want to use
    client := gobalt.NewClientWithAPI("http://localhost:9000")

    // Then simply get the media
    media, err := gobalt.Post(context.Background(),gobalt.PostParams{URL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ"})
    if err != nil {
        panic(err)
    }

    // Then simply stream it to save it or to send it somewhere else
    stream, err := media.Stream(context.Background())
    if err != nil {
        panic(err)
    }
    // You must close the stream once you are done
    defer stream.Close()

    file, err := os.OpenFile(media.Filename, os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        panic(err)
    }

    // Read and write into the file into EOF from the stream
    _, err = file.ReadFrom(stream)
    if err != nil {
        panic(err)
    }
}
```

## Private Instance is Required

Since the removal of the api.cobalt.tools public v7 API, a private hosted
instance is required to use with this library.

## Support cobalt.tools

The best way to support this project is by supporting
[cobalt.tools](https://cobalt.tools/),
its an amazing FREE, AD FREE, NO TRACKING service, which makes life easier. So
go there and donate, if possible.
