package storage

import "time"

type User struct {
	Id           string //`json:"id"`
	Username     string
	Email        string
	Password     string
	LastOnline   time.Time
	OnlineStatus bool
}

func (usr User) getUser() {}

func (user User) setUserState() {}
