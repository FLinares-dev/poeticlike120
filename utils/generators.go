package utils

import (
	"math/rand"
	"time"
)

type UserInfo struct {
	Name     string
	Surname  string
	Email    string
	Password string
}

func New(idx int) UserInfo {
	rand.Seed(time.Now().Unix())
	name := names[rand.Intn(len(names))]
	surname := surnames[rand.Intn(len(surnames))]
	separator := separators[rand.Intn(len(separators))]

	userInfo := UserInfo{
		Name:     name,
		Surname:  surname,
		Email:    surname + separator + name + "@mail.com",
		Password: GenericPWD,
	}

	return userInfo
}
