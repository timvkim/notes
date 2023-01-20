package service

import "github.com/timvkim/notes/internal/repository"

type Service struct {
	repo *repository.Repo
}

func NewService(repo *repository.Repo) *Service {
	return &Service{
		repo: repo,
	}
}
