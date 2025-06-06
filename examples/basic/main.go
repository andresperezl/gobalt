package main

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/andresperezl/gobalt/v2"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need to provide an URL as argument")
	}

	client := gobalt.NewCobaltWithAPI("http://localhost:9000")
	resp, err := client.Post(context.Background(), gobalt.PostRequest{
		URL: os.Args[1],
	})
	if err != nil {
		log.Fatal(err)
	}
	var stream io.ReadCloser

	if resp.Status == gobalt.ResponseStatusPicker {
		for _, pi := range resp.Picker {
			if pi.Type == gobalt.PickerItemTypeVideo {
				stream, err = pi.Stream(context.Background())
			}
		}
	} else {
		stream, err = resp.Stream(context.Background())
	}
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	if _, err := os.Stdout.ReadFrom(stream); err != nil {
		log.Fatal(err)
	}
}
