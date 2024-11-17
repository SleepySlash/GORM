package model

import (
	"gorm.io/gorm"
)

type Account interface {
	Save(person Person) error
	FindName(name string) (Person, error)
	FindAll() ([]Person, error)
	Delete(name string) (Person, error)
}

type accountDB struct {
	DB *gorm.DB
}

func CreateConntection(db *gorm.DB) Account {
	return &accountDB{
		DB: db,
	}
}

func (p *accountDB) Save(person Person) error {
	result := p.DB.Create(&person)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *accountDB) FindName(name string) (Person, error) {
	var person Person
	res := p.DB.Where("name = ?", name).First(&person)
	if res.Error != nil {
		return person, res.Error
	}
	return person, nil
}

func (p *accountDB) FindAll() ([]Person, error) {
	var person []Person
	res := p.DB.Find(&person)
	if res.Error != nil {
		return nil, res.Error
	}
	return person, nil
}

func (p *accountDB) Delete(name string) (Person, error) {
	var person Person
	res := p.DB.Where("name = ?", name).First(&person)
	if res.Error != nil {
		return person, res.Error
	}
	res = p.DB.Where("name = ?", name).Delete(person)
	if res.Error != nil {
		return person, res.Error
	}
	return person, nil
}
