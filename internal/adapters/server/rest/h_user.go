package rest

import (
	"net/http"
	"strconv"

	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func (o *St) hUserCreate(ctx *gin.Context) {
	reqObj := &entities.UserCUSt{}

	result := o.ucs.UserCreate(ctx, reqObj)

	ctx.JSON(http.StatusOK, result)
}

func (o *St) hUserGet(ctx *gin.Context) {
	id := ctx.Param("id")
	userid, _ := strconv.Atoi(id)

	result, err := o.ucs.UserGet(ctx, userid)
	if err != nil {
		o.lg.Errorw("error", err)
	}

	ctx.JSON(http.StatusOK, result)
}
