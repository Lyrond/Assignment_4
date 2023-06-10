package usecase

import (
	"Assignment_4/services/contact/internal/domain"
	"Assignment_4/services/contact/internal/repository"
)

type UseCase interface {
	CreateContact(user *domain.Contact) (int, error)
	ReadContact(id int) (*domain.Contact, error)
	UpdateContact(id int) error
	DeleteContact(id int) error
	CreateGroup(group *domain.Group) (int, error)
	ReadGroup(id int) (*domain.Group, error)
	AddContact(id int) error
}

type useCase struct {
	repo repository.Repo
}

func New(repo repository.Repo) *useCase {
	return &useCase{repo: repo}
}

func (u *useCase) CreateContact(user *domain.Contact) (int, error) {
	return 0, nil
}

func (u *useCase) ReadContact(id int) (*domain.Contact, error) {
	return nil, nil
}

func (u *useCase) UpdateContact(id int) error {
	return nil
}

func (u *useCase) DeleteContact(id int) error {
	return nil
}

func (u *useCase) CreateGroup(group *domain.Group) (int, error) {
	return 0, nil
}

func (u *useCase) ReadGroup(id int) (*domain.Group, error) {
	return nil, nil
}

func (u *useCase) AddContact(id int) error {
	return nil
}
