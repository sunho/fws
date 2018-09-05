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
	ID            int    `json:"id" xorm:"pk autoincr"`
	Name          string `json:"name"`
	WebhookSecret string `json:"webhook_secret"`
	GitURL        string `json:"git_url"`
}

type UserBot struct {
	UserID int `xorm:"pk"`
	BotID  int `xorm:"pk"`
}

type Webhook struct {
	Hash   string `xorm:"pk"`
	Secret string
	BotID  int
}

type Volume struct {
	BotID int    `json:"bot_id" xorm:"pk"`
	Name  string `json:"name" xorm:"pk"`
	Size  int64  `json:"size"`
	Path  string `json:"path"`
}

type Config struct {
	BotID int    `json:"bot_id" xorm:"pk"`
	Name  string `json:"name" xorm:"pk"`
	Path  string `json:"path"`
	Value string `json:"value"`
}

type Env struct {
	BotID int    `json:"bot_id" xorm:"pk"`
	Name  string `json:"name" xorm:"pk"`
	Value string `json:"value"`
}

type Build struct {
	BotID   int       `json:"bot_id" xorm:"pk"`
	Number  int       `json:"number" xorm:"pk"`
	Success bool      `json:"success"`
	Created time.Time `json:"created"`
}

type BuildLog struct {
	BotID  int `xorm:"pk"`
	Number int `xorm:"pk"`
	Logged []byte
}
