package main

import "time"

// Embedding the User struct anonymously
type AdminAnonUser struct {
	email    string
	password string
	User
}

// embedding the User struct with custom name
type Admin struct {
	email    string
	password string
	user     User
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		email:    email,
		password: password,
		user: User{
			"Admin",
			"admin",
			"01/01/2001",
			time.Now(),
		},
	}
}

func NewAdminAnon(email, password string) *AdminAnonUser {
	return &AdminAnonUser{
		email:    email,
		password: password,
		User: User{
			"Admin",
			"admin",
			"01/01/2001",
			time.Now(),
		},
	}
}
