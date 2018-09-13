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

type RunStatus interface {
	sealed()
}

type RunStatusRunning struct {
	Type    string `json:"type"`
	Seconds int    `json:"seconds"`
}

func (RunStatusRunning) sealed() {}

type RunStatusFailed struct {
	Type    string `json:"type"`
	Seconds int    `json:"seconds"`
}

func (RunStatusFailed) sealed() {}

type RunStatusPending struct {
	Type string `json:"type"`
}

func (RunStatusPending) sealed() {}

type RunBot struct {
	*Bot
	Configs []*BotConfig
	Volumes []*BotVolume
	Envs    []*BotEnv
}
