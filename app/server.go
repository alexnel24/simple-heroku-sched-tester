package app

import "net/http"

type Server struct{
	router *http.ServeMux
}

func NewServer() *Server{
	s := &Server{
		http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}