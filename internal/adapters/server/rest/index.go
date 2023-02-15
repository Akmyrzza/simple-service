package rest

import (
	"net/http"

	"github.com/Akmyrzza/simple-service/internal/adapters/logger"
	"github.com/Akmyrzza/simple-service/internal/domain/usecases"
	"github.com/gin-gonic/gin"
)

type St struct {
	lg  logger.Lite
	ucs *usecases.St
}

func GetHandler(lg logger.Lite, ucs *usecases.St) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	s := &St{lg: lg, ucs: ucs}

	r.POST("/ping", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })
	r.POST("/user", s.hUserCreate)

	return r
}
