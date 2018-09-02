package model

type BuildStatus struct {
	Running bool
	Step    string
}

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
