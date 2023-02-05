package core

import "github.com/Akmyrzza/simple-service/internal/adapters/repo"

type St struct {
	repo repo.Repo

	User *User
}

func New(repo repo.Repo) *St {
	c := &St{
		repo: repo,
	}

	c.User = NewUser(c)

	return c
}