package model

type UserDomainInterface interface {
	EncryptPassword()
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() string
	SetID(string)
}

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}
