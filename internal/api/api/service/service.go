package service

import (
	"test/internal/api/store"
)

type Service interface {
}

type service struct {
	factory store.Factory
}

func NewService(factory store.Factory) Service {
	return &service{
		factory: factory,
	}
}
