package server

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"time"
	"uniovi-localizeme/internal/core/ports/impl"
	"uniovi-localizeme/tools"
)

type Server struct {
	server *http.Server
}

func CreateServer(port string) *Server {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	config := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(config.Handler)
	initControllers(r)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &Server{server}
}

func (s *Server) Shutdown() error {
	log.Printf("Stopping on http://localhost%s", s.server.Addr)
	return s.server.Shutdown(context.TODO())
}

func (s *Server) Start() {
	log.Printf("Server running on http://localhost%s", s.server.Addr)
	log.Printf("%s", s.server.ListenAndServe())
}

func initControllers(r *chi.Mux) {
	log.Printf("%s: start", tools.GetCurrentFuncName())
	userPort := impl.CreateUserPort()
	stagePort := impl.CreateStagePort()
	languagePort := impl.CreateLanguagePort()
	groupPort := impl.CreateGroupPort()
	baseStringPort := impl.CreateBaseStringPort()
	xliffPort := impl.CreateXliffPort()
	userPort.InitRoutes(r)
	stagePort.InitRoutes(r)
	languagePort.InitRoutes(r)
	groupPort.InitRoutes(r)
	baseStringPort.InitRoutes(r)
	xliffPort.InitRoutes(r)
	log.Printf("%s: end", tools.GetCurrentFuncName())
}
