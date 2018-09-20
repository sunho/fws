package basic

import (
	"errors"
	"io"
	"time"

	"github.com/sunho/fws/server/model"
	"github.com/sunho/fws/server/runtime"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	ErrAlreadyExist = errors.New("basic: already exist")
)

type Runner struct {
	cli *kubernetes.Clientset

	Namespace     string
	volumeManager *VolumeManager
}

func NewRunnerFromFile(namespace string, path string, nfspath string, nfsaddr string) (*Runner, error) {
	conf, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		return nil, err
	}

	cli, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	vol := &VolumeManager{
		NfsPath: nfspath,
		NfsAddr: nfsaddr,
	}

	return &Runner{
		cli:           cli,
		volumeManager: vol,
		Namespace:     namespace,
	}, nil
}

func NewRunnerFromCluster(namespace string, nfspath string, nfsaddr string) (*Runner, error) {
	conf, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return nil, err
	}

	cli, err := kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	vol := &VolumeManager{
		NfsPath: nfspath,
		NfsAddr: nfsaddr,
	}

	return &Runner{
		cli:           cli,
		volumeManager: vol,
		Namespace:     namespace,
	}, nil
}

func (r *Runner) Exists(bot *model.RunBot) (bool, error) {
	_, err := r.getDeployment(bot.Bot)
	if err == nil {
		return false, err
	}
	if err != ErrNotExists {
		return true, nil
	}

	confs, err := r.getConfigs(bot.Bot)
	if err != nil {
		return false, err
	}
	if len(confs) != 0 {
		return true, nil
	}

	return false, nil
}

func (r *Runner) Put(bot *model.RunBot) error {
	err := r.deleteDeployment(bot.Bot)
	if err != nil {
		return err
	}

	err = r.deleteConfigs(bot.Bot)
	if err != nil {
		return err
	}

	err = r.volumeManager.Init(bot.ID)
	if err != nil {
		return err
	}

	names, err := r.volumeManager.List(bot.ID)
	if err != nil {
		return err
	}

L:
	for _, n := range names {
		for _, v := range bot.Volumes {
			if v.Name == n {
				continue L
			}
		}

		err = r.volumeManager.Delete(bot.ID, n)
		if err != nil {
			return err
		}
	}

L2:
	for _, v := range bot.Volumes {
		for _, n := range names {
			if v.Name == n {
				continue L2
			}
		}
		err = r.volumeManager.Create(bot.ID, v.Name)
		if err != nil {
			return err
		}
	}

	deployCli := r.cli.Apps().Deployments(r.Namespace)
	configCli := r.cli.Core().ConfigMaps(r.Namespace)

	confs := r.makeConfigs(bot)
	// potential bug
	for _, conf := range confs {
		_, err = configCli.Create(&conf)
		if err != nil {
			return err
		}
	}

	deploy := r.makeDeployment(bot)
	_, err = deployCli.Create(deploy)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) Delete(bot *model.Bot) error {
	err := r.deleteDeployment(bot)
	if err != nil {
		return err
	}

	err = r.deleteConfigs(bot)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) UpdateBuild(bot *model.Bot) error {
	deploy, err := r.getDeployment(bot)
	if err != nil {
		return err
	}

	deployCli := r.cli.Apps().Deployments(r.Namespace)
	deploy.Spec.Template.Spec.Containers[0].Image = bot.BuildResult
	_, err = deployCli.Update(deploy)
	return err
}

func (r *Runner) Run(bot *model.Bot) error {
	return nil
}

func (r *Runner) Restart(bot *model.Bot) error {
	return r.deletePods(bot)
}

func (r *Runner) Stop(bot *model.Bot) error {
	return nil
}

func (r *Runner) Status(bot *model.Bot) (model.RunStatus, error) {
	pods, err := r.getPods(bot)
	if err != nil {
		return nil, err
	}

	if len(pods) == 0 {
		return nil, runtime.ErrNotExists
	}

	if len(pods) != 1 {
		return &model.RunStatusPending{
			Type: "pending",
		}, nil
	}

	p := pods[0]
	var ready *apiv1.PodCondition
	for _, c := range p.Status.Conditions {
		if c.Type == apiv1.PodReady {
			ready = &c
			break
		}
	}
	if ready == nil {
		return &model.RunStatusPending{
			Type: "pending",
		}, nil
	}

	if ready.Status == apiv1.ConditionTrue {
		return &model.RunStatusRunning{
			Type:    "running",
			Seconds: int(time.Now().Sub(ready.LastTransitionTime.Time).Seconds()),
		}, nil
	}
	return &model.RunStatusFailed{
		Type:    "failed",
		Seconds: int(time.Now().Sub(ready.LastTransitionTime.Time).Seconds()),
	}, nil
}

func (r *Runner) Log(bot *model.Bot) ([]byte, error) {
	return nil, nil
}

func (r *Runner) DownloadVolume(volume *model.BotVolume) (io.Reader, error) {
	return nil, nil
}

func (r *Runner) VolumeUsed(volume *model.BotVolume) (int64, error) {
	return 0, nil
}
