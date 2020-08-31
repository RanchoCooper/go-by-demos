package service

import (
	"database/sql"
	"log"

	"github.com/google/wire"
)

// IUserRepo
type IUserRepo interface {
	AddUser()
}

// UserRepo implement the IUserRepo interface
type UserRepo struct {
	*sql.DB
}

// AddUser
func (u *UserRepo) AddUser() {
	log.Println("add user")
}

// NewUserRepo
func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{}
}

var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(IUserRepo), new(*UserRepo)))
