package users

import "muzucode/fawn/users-passwords/passwords"

type User struct {
	Id       int
	Name     string
	Theme    Theme
	Password passwords.Password
}

type Theme struct {
	Id              int
	Name            string
	PrimaryColor    string
	SecondaryColor  string
	TertiaryColor   string
	QuaternaryColor string
}
