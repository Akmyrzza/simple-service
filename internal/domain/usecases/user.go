package usecases

import (
	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

func (u *St) UserCreate(ctx *gin.Context, obj *entities.UserCUSt) string {

	result := u.cr.User.Create(ctx, obj)

	return result
}

func (u *St) UserGet(ctx *gin.Context, id int) (*entities.UserSt, error) {

	result, err := u.cr.User.Get(ctx, id)

	return result, err
}
