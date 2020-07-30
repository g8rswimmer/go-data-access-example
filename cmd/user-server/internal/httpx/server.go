package httpx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/g8rswimmer/go-data-access-example/pkg/api/response"
	"github.com/gorilla/mux"
)

// Router will add the routes to the mux
type Router interface {
	Add(*mux.Router)
}

// Info contains the information on the service
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Server handles the http server for the service
type Server struct {
	svr  *http.Server
	info Info
}

const shutdownTO = time.Second * 10

// NewServer creates a new server
func NewServer(info Info, port string, rto, wto time.Duration) *Server {
	return &Server{
		svr: &http.Server{
			Addr:         fmt.Sprintf(":%s", port),
			ReadTimeout:  rto,
			WriteTimeout: wto,
		},
		info: info,
	}
}

// Start the server with the routes
func (s *Server) Start(routers []Router) {

	s.svr.Handler = s.handler(routers)

	go func() {
		if err := s.svr.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) == false {
			panic(err)
		}
	}()
}

func (s Server) handler(routers []Router) http.Handler {
	r := mux.NewRouter().StrictSlash(true)
	r.Methods(http.MethodGet).Path("/").Handler(s.index()).Name("info")

	apis := r.PathPrefix("/v1").Subrouter()
	for _, router := range routers {
		router.Add(apis)
	}
	return r
}

// Shutdown will gracefuly shutdown the server
func (s *Server) Shutdown(ctx context.Context) error {
	ctxTO, cancel := context.WithTimeout(ctx, shutdownTO)
	defer cancel()
	return s.svr.Shutdown(ctxTO)
}

func (s *Server) index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response.JSON(w, http.StatusOK, s.info)
	}
}
