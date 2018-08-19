package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Passhash string `json:"passhash"`
}

type UserInvite struct {
	Username string `json:"username"`
	Key      string `json:"key"`
}

type Bot struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	GitURL string `json:"git_url"`
}

type Webhook struct {
	Hash   string
	BotID  int
	GitURL string
}

type Volume struct {
	BotID int
	Name  string
	Size  int64
	Path  string
}

type Config struct {
	BotID int
	Name  string
	Path  string
}

type Env struct {
	BotID int
	Name  string
}

type Build struct {
	Number  int
	BotID   int
	Success bool
	Created time.Time
}
