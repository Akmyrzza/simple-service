package core

import (
	"github.com/Akmyrzza/simple-service/internal/domain/entities"
	"github.com/gin-gonic/gin"
)

type User struct {
	r *St
}

func NewUser(r *St) *User {
	return &User{r: r}
}

func (c *User) Create(ctx *gin.Context, obj *entities.UserCUSt) string {
	result := c.r.repo.UserCreate(ctx, obj)

	return result
}
