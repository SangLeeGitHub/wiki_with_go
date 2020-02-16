package main

import "net/http"

type Server struct {

	*router
	startHandler HandlerFunc
}

func NewServer() *Server {

	r := &router{make(map[string]map[string]HandlerFunc)}
	s := &Server{router: r}

	return s
}

func (s *Server) Run(addr string) {
	// start Handler
	s.startHandler = s.router.handler()

	// Start Web Server
	if err := http.ListenAndServe(addr, s); err != nil {
		panic(err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := &Context{Params: make(map[string]interface{}), ResponseWriter: w, Request: r}
	for k, v := range r.URL.Query() {
		c.Params[k] = v[0]
	}

	s.startHandler(c)
}