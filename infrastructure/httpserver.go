package infrastructure

import (
	"net/http"
)

type HTTPServer struct {
	mux  *http.ServeMux
	port string
}

func NewHTTPServer(port string, webservices map[string]http.Handler, loggerMW interface{}) *HTTPServer {
	server := &HTTPServer{
		mux:  http.NewServeMux(),
		port: port,
	}
	server.setRoutes(webservices, loggerMW)
	return server
}

func (h *HTTPServer) Serve() error {
	return http.ListenAndServe(h.port, h.mux)
}

func (h *HTTPServer) setRoutes(webservices map[string]http.Handler, loggerMW interface{}) {
	// h.mux.Handle(people, loggerMW.Log(webservices["PersonService"]))
}
