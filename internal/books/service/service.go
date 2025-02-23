package booksservice

import (
	"github.com/sunzhqr/book_library/internal/books/model"
)

type Repository interface {
	Get() ([]model.Book, error)
	Add(req model.Book) error
}

type Wheather interface {
	GetWeather()
}

type Service struct {
	repo     Repository
	wheather Wheather
}

func NewService(repo Repository, wheather Wheather) *Service {
	return &Service{
		repo:     repo,
		wheather: wheather,
	}
}

func (s *Service) Get() ([]model.Book, error) {
	s.wheather.GetWeather()
	return s.repo.Get()
}

func (s *Service) Add(req model.Book) error {
	return s.repo.Add(req)
}
