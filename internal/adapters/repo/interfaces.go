package repo

import (
	"context"

	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type Repo interface {
	UserCreate(ctx *gin.Context, obj *entities.UserCUSt) string
	UserGet(ctx context.Context, id int) (*entities.UserSt, error)
}
