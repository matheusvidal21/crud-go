package model

type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}

func (ud *userDomain) GetPassword() string {
	return ud.Password
}

func (ud *userDomain) GetName() string {
	return ud.Name
}

func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

func (ud *userDomain) GetID() string {
	return ud.ID
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}
