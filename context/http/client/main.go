package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	if err != nil {
		log.Fatal(err)
	}

	ctx = context.WithValue(ctx, "foo", "bar") // the value do not receive from the context of server

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)

	defer cancel()

	req = req.WithContext(ctx)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	io.Copy(os.Stdout, res.Body)
}
