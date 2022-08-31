package httprest

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joaocprofile/goh/core"
	"github.com/joaocprofile/goh/database/cachedb"
	"github.com/joaocprofile/goh/database/sqlg"
	env "github.com/joaocprofile/goh/environment"
	"github.com/joaocprofile/goh/security"
)

type Server struct {
	router     *mux.Router
	httpServer *http.Server
	routes     []core.Route
}

func NewHttpRest() *Server {
	router := mux.NewRouter()
	httpsrv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("127.0.0.1:%d", env.Get().APIPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	srv := Server{
		httpServer: httpsrv,
		router:     router,
		routes:     nil,
	}
	return &srv
}

func (server *Server) AddDatabase() *Server {
	conn := sqlg.NewConnection()
	log.Println(core.Yellow("Connected to the database " + env.Get().DBConnection.DBDriver))
	defer conn.Close()
	return server
}

func (server *Server) AddCache() *Server {
	cachedb := cachedb.NewConnection()
	log.Println(core.Yellow("Connected to the cache " + env.Get().CacheDB.CacheDriver))
	defer cachedb.Close()
	return server
}

func (server *Server) AddRoutes(routes []core.Route) {
	server.routes = append(server.routes, routes...)
	for _, route := range server.routes {
		if route.Authentication {
			server.router.HandleFunc(
				route.URI,
				security.EnsureAuth(route.Function),
			).Methods(route.Method)
		} else {
			server.router.HandleFunc(
				route.URI,
				route.Function,
			).Methods(route.Method)
		}
	}
}

func (server *Server) ListenAndServe() {
	log.Println(core.Yellow("Service running on the port"), env.Get().APIPort)
	log.Fatal(server.httpServer.ListenAndServe())
}
