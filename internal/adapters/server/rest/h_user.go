package rest

import (
	"net/http"

	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func (o *St) hUserCreate(ctx *gin.Context) {
	reqObj := &entities.UserCUSt{}

	result := o.ucs.UserCreate(ctx, reqObj)

	ctx.JSON(http.StatusOK, result)
}
