package handler

import (
	"compress/gzip"
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go-pgdb/config"
	"go-pgdb/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Env struct {
	Router *mux.Router
	Config *config.Configuration
	PGDB   db.CRUDOperations
}

func HandleRequests() {
	log.Println("Handling Requests.!")
	var wait time.Duration
	env := &Env{mux.NewRouter(), nil, nil}
	env.SetupEnvironment()
	env.SetupRoutes()

	conn, err := db.Init(env.Config)
	if err != nil {
		log.Panic(err)
	}
	env.PGDB = conn

	srv := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 100,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: handlers.CombinedLoggingHandler(
			os.Stdout,
			handlers.CompressHandlerLevel(
				authMiddleware(env.Router),
				gzip.DefaultCompression),
		),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func authMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

func (env *Env) SetupRoutes() {
	subrouter := env.Router.StrictSlash(true).PathPrefix("/api").Subrouter()
	subrouter.HandleFunc("/forum/all", env.GetAllForums)
	subrouter.HandleFunc("/forum", env.CreateForum).Methods("POST")
	subrouter.HandleFunc("/forum", env.UpdateForum).Methods("PUT")
	subrouter.HandleFunc("/forum/{id}", env.GetForumById)
	subrouter.HandleFunc("/thread/forumId/{id}", env.GetThreadsByForumId)
	subrouter.HandleFunc("/post/threadId/{id}", env.GetPostsByThreadId)
	subrouter.HandleFunc("/post/forumId/{id}", env.GetPostsByForumId)
}