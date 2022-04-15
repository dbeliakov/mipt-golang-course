package middleware

import (
	"log"
	"net/http"
)

func Recover(logger *log.Logger) func(http.Handler) http.Handler {
	return nil
}
