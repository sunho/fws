package model

type BuildStatus int

const (
	BuildStatusBuilding BuildStatus = iota
	BuildStatusNone
)

type RunStatus int

const (
	RunStatusRunning RunStatus = iota
	RunStatusStopped
	RunStatusCrash
)

type RunBot struct {
	*Bot
	Configs []*Config
	Volumes []*Volume
	Envs    []*Env
}
