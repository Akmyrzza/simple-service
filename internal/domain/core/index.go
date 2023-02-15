package core

import (
	"github.com/Akmyrzza/simple-service/internal/adapters/logger"
	"github.com/Akmyrzza/simple-service/internal/adapters/repo"
)

type St struct {
	lg   logger.Lite
	repo repo.Repo

	User *User
}

func New(lg logger.Lite, repo repo.Repo) *St {
	c := &St{
		lg:   lg,
		repo: repo,
	}

	c.User = NewUser(c)

	return c
}
