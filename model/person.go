package model

type Person struct {
	Name   string `gorm:"type:varchar(225);primary_key"`
	Age    int    `gorm:"type:int"`
	Number string `gorm:"type:varchar(225)"`
}

func NewPerson(name, number string, age int) Person {
	return Person{
		Name:   name,
		Age:    age,
		Number: number,
	}
}

func (Person) TableName() string {
	return "accounts"
}
