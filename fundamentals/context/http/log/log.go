package log

import (
	"context"
	"log"
	"math/rand"
	"net/http"
)

type key int

const requestIdKey = key(42)

func Println(ctx context.Context, msg string) {
	id, ok := (ctx.Value(requestIdKey)).(int64)

	if !ok {
		log.Fatal("Do not have any id to log")
		return
	}

	log.Printf("[%d]: %v", id, msg)
}

func Decorator(f http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		randNum := rand.Int63()

		ctx = context.WithValue(ctx, requestIdKey, randNum)

		f(w, r.WithContext(ctx))
	}
}
