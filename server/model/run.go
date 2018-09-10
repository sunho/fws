package model

import "time"

type BuildStatus interface {
	sealed()
}

type BuildStatusBuilt struct {
	Type    string    `json:"type"`
	Success bool      `json:"success"`
	Number  int       `json:"number"`
	Created time.Time `json:"created"`
}

func (BuildStatusBuilt) sealed() {}

type BuildStatusBuilding struct {
	Type    string `json:"type"`
	Pending bool   `json:"pending"`
	Step    string `json:"step"`
}

func (BuildStatusBuilding) sealed() {}

type RunStatus int

const (
	RunStatusRunning RunStatus = iota
	RunStatusStopped
	RunStatusCrash
)

type RunBot struct {
	*Bot
	Configs []*BotConfig
	Volumes []*BotVolume
	Envs    []*BotEnv
}
