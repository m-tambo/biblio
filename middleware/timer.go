package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ExecutionTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r) // Call the next handler
		fmt.Printf("Execution time: %s \n", time.Now().Sub(t).String())
	})
}