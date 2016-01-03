package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	Name        string
	HoursActive [24]int

	Actions  []string
	Messages []string

	JoinCount int
	PartCount int
}

func (u *User) RandomQuote() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return u.Messages[r.Intn(len(u.Messages))]
}

func (u *User) String() string {
	str := u.Name + ": "
	str += strconv.Itoa(len(u.Actions)) + " actions, "
	str += strconv.Itoa(len(u.Messages)) + " messages, "
	str += fmt.Sprint(u.HoursActive)
	return str
}

func NewUser(nick string) *User {
	return &User{
		Name: nick,
	}
}
