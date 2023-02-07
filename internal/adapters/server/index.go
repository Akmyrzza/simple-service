package server

import (
	"net/http"
	"time"
)

const (
	ReadHeaderTimeout = 10 * time.Second
	ReadTimeout       = 2 * time.Minute
	MaxHeaderBytes    = 300 * 1024
)

type St struct {
	addr   string
	server *http.Server
}

func Start(addr string, handler http.Handler) *St {
	s := &St{
		addr: addr,
		server: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadHeaderTimeout: ReadHeaderTimeout,
			ReadTimeout:       ReadTimeout,
			MaxHeaderBytes:    MaxHeaderBytes,
		},
	}

	//go func() {
	//err :=
	s.server.ListenAndServe()
	//if err != nil {
	//fmt.Print("http server closed", err)
	//}
	//}()

	return s
}
