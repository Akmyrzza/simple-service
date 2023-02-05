package server

import (
	"time"
)

const (
	ReadHeaderTimeout = 10 * time.Second
	ReadTimeout       = 2 * time.Minute
	MaxHeaderBytes    = 300 * 1024
)

func Start(addr string) {

}

