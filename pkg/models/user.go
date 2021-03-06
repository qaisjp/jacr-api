package models

import "time"

type User struct {
	ID int `db:"id"`

	Username  string    `db:"username" valid:"stringlength(1|255),required"`
	Password  string    `db:"password" valid:"stringlength(5|100),required"`
	Email     string    `db:"email" valid:"email,stringlength(1|254),required"`
	Slug      string    `db:"slug" valid:"stringlength(1|255),required"`
	Level     int       `db:"level"`
	Banned    bool      `db:"banned"`
	Activated bool      `db:"activated"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
