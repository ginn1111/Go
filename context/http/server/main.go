package main

import (
	"context"
	"fmt"
	"go/context/http/log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", log.Decorator(handler))

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// override the key of log func in log package
	// [WORK AROUND]: defined package's scope type
	ctx = context.WithValue(ctx, int(42), int64(111))

	log.Println(ctx, "handler start")
	defer log.Println(ctx, "handler end")

	fmt.Println(ctx.Value("foo"))

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "Hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
