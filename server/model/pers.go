package model

import "time"

type User struct {
	ID       int    `json:"id" xorm:"pk autoincr"`
	Admin    bool   `json:"admin"`
	Version  int    `xorm:"version"`
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
	Version       int    `json:"version" xorm:"version"`
	BuildResult   string `json:"build_result"`
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

type BotVolume struct {
	BotID   int    `json:"bot_id" xorm:"pk"`
	Name    string `json:"name" xorm:"pk"`
	Version int    `json:"version" xorm:"version"`
	Size    int64  `json:"size"`
	Path    string `json:"path"`
}

type BotConfig struct {
	BotID   int    `json:"bot_id" xorm:"pk"`
	Name    string `json:"name" xorm:"pk"`
	Version int    `json:"version" xorm:"version"`
	Path    string `json:"path"`
	File    string `json:"file"`
	Value   string `json:"value"`
}

type BotEnv struct {
	BotID   int    `json:"bot_id" xorm:"pk"`
	Name    string `json:"name" xorm:"pk"`
	Version int    `json:"version" xorm:"version"`
	Value   string `json:"value"`
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
