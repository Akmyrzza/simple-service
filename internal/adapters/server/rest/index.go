package rest

import (
	"net/http"

	"github.com/Akmyrzza/simple-service/internal/domain/usecases"
	"github.com/gin-gonic/gin"
)

type St struct {
	ucs *usecases.St
}

func GetHandler(ucs *usecases.St) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	s := &St{ucs: ucs}

	//r.POST("/ping", sqlite.Ping)
	r.POST("/user", s.hUserCreate)

	return r
}
