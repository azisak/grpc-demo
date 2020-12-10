package server

import "errors"

var ErrNotFound = errors.New("not found")

type User struct {
	ID   int64
	Name string
}

type Service interface {
	GetUser(id int64) (User, error)
}

type svc struct {
	m map[int64]User
}

func NewService() Service {
	return &svc{
		m: map[int64]User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
			3: {ID: 3, Name: "Carol"},
		},
	}
}

func (s *svc) GetUser(id int64) (result User, err error) {
	if result, ok := s.m[id]; ok {
		return result, nil
	}
	return result, ErrNotFound
}
