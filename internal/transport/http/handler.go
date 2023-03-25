package http

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
)

// Handler - stores pointer to our user service
type Handler struct {
	Router  chi.Router
	Service UserService
	Server  *http.Server
}

// Response object
type Response struct {
	Message string `json:"message"`
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service UserService) *Handler {
	log.Info("setting up our handler")

	h := &Handler{
		Service: service,
	}
	h.Router = chi.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	// return our wonderful handler
	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.Get("/alive", h.AliveCheck)
	h.Router.Get("/api/user/{email}", h.GetUser)
	h.Router.Post("/api/user", h.CreateUser)
	h.Router.Put("/api/user/{email}", h.UpdateUser)
	h.Router.Delete("/api/user/{email}", h.DeleteUser)

}

func (h *Handler) AliveCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
		panic(err)
	}
}

// Serve - gracefully serves our newly set up handler function
func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Create a deadline to wait for
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutting down gracefully")
	return nil
}
