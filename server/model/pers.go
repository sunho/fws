package model

import "time"

type User struct {
	ID       int    `json:"id" xorm:"pk autoincr"`
	Admin    bool   `json:"admin"`
	Username string `json:"username" xorm:"unique"`
	Nickname string `json:"nickname" xorm:"unique"`
	Passhash string `json:"passhash"`
}

type UserInvite struct {
	Username string `json:"username" xorm:"pk"`
	Admin    bool   `json:"admin"`
	Key      string `json:"key"`
}

type Bot struct {
	ID     int    `json:"id" xorm:"pk autoincr"`
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
	BotID   int `xorm:"pk"`
	Number  int `xorm:"pk"`
	Success bool
	Created time.Time
}

type BuildLog struct {
	BotID  int `xorm:"pk"`
	Number int `xorm:"pk"`
	Logged []byte
}
