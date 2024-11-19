package services

import "gorm/postgresql/model"

type Service interface {
	AddNewAccount(name, number string, age int) error
	FindAccount(name string) (model.Person, error)
	FindAllAccounts() ([]model.Person, error)
	DeleteAccount(name string) (model.Person, error)
}

type service struct {
	account model.Account
}

func NewService(gorm model.Account) Service {
	return &service{
		account: gorm,
	}
}

func (s *service) AddNewAccount(name, number string, age int) error {
	person := model.NewPerson(name, number, age)
	err := s.account.Save(person)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) FindAllAccounts() ([]model.Person, error) {
	person, err := s.account.FindAll()
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (s *service) FindAccount(name string) (model.Person, error) {
	person, err := s.account.FindName(name)
	if err != nil {
		return person, err
	}
	return person, nil
}

func (s *service) DeleteAccount(name string) (model.Person, error) {
	var person model.Person
	person, err := s.account.Delete(name)
	if err != nil {
		return person, err
	}
	return person, nil
}
