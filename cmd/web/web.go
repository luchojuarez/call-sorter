package web

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luchojuarez/call-sorter/cmd/container"
	"github.com/luchojuarez/call-sorter/internal/domain/invoice"
	"github.com/luchojuarez/call-sorter/internal/handlers"
)

func NewWebServer() *http.Server {
	c := container.GetSimpleContainer()

	proccesor := invoice.NewProcessor(c.GetLocalCallRepository(), c.GetRestUserRepository())
	handler := handlers.NewHandler(proccesor)

	router := mux.NewRouter()

	router.Use(setJsonResponse)

	router.HandleFunc("/v1/invoice/{year}/{month}", handler.Generate)

	return &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func setJsonResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
