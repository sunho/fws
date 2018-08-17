package model

import "time"

type User struct {
	ID       int
	Username string
	Passhash string
}

type Bot struct {
	ID     int
	Name   string
	GitURL string
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
