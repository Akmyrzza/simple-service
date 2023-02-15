package server

import (
	"net/http"
	"time"

	"github.com/Akmyrzza/simple-service/internal/adapters/logger"
)

const (
	ReadHeaderTimeout = 10 * time.Second
	ReadTimeout       = 2 * time.Minute
	MaxHeaderBytes    = 300 * 1024
)

type St struct {
	lg logger.Lite

	addr   string
	server *http.Server
}

func Start(lg logger.Lite, addr string, handler http.Handler) *St {
	s := &St{
		lg:   lg,
		addr: addr,
		server: &http.Server{
			Addr:              addr,
			Handler:           handler,
			ReadHeaderTimeout: ReadHeaderTimeout,
			ReadTimeout:       ReadTimeout,
			MaxHeaderBytes:    MaxHeaderBytes,
		},
	}

	s.lg.Infow("Start rest-api", "addr", s.server.Addr)
	//go func() {
	//err :=
	s.server.ListenAndServe()
	//if err != nil {
	//fmt.Print("http server closed", err)
	//}
	//}()

	return s
}
