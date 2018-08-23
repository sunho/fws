package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Admin    bool   `json:"admin"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
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
	BotID int    `json:"bot_id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Path  string `json:"path"`
}

type Config struct {
	BotID int    `json:"bot_id"`
	Name  string `json:"name"`
	Path  string `json:"path"`
}

type Env struct {
	BotID int    `json:"bot_id"`
	Name  string `json:"name"`
}

type Build struct {
	Number  int
	BotID   int
	Success bool
	Created time.Time
}
