package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/info", func(rw http.ResponseWriter, req *http.Request) {
		uid, err := UserIDFromContext(req.Context())
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Req")
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte(fmt.Sprintf("Hello, %d", uid)))
	})

	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatal(err)
	}
}

func Auth(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		authValue := req.Header.Get("Authorization")
		// Parse token and auth
		userID, err := strconv.Atoi(authValue)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(rw, req.WithContext(ContextWithUserID(req.Context(), userID)))
	}
	return http.HandlerFunc(fn)
}

type userIDKey struct{}

func ContextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, userIDKey{}, userID)
}

func UserIDFromContext(ctx context.Context) (int, error) {
	uid := ctx.Value(userIDKey{})
	if uid == nil {
		return 0, fmt.Errorf("no user id in context")
	}
	return uid.(int), nil
}
