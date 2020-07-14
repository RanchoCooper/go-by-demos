package user

import (
	"github.com/RanchoCooper/go-by-demos/third-party-library/gomock/person"
)

type User struct {
	Person person.Male
}

func NewUser(p person.Male) *User {
	return &User{Person: p}
}

func (u *User) GetUserInfo(id int64) error {
	return u.Person.Get(id)
}