package basic

import (
	"errors"
	"path/filepath"
	"strconv"

	"github.com/sunho/fws/server/model"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ErrMultipleDeployments = errors.New("basic: mutilple deployments")
	ErrNotExists           = errors.New("basic: doesn't exist")
)

const (
	labelIDKey     = "fws-bot-id"
	labelConfigKey = "fws-bot-config"
	configKey      = "config"
)

func int32Ptr(i int32) *int32 { return &i }

func (r *Runner) getConfigs(bot *model.RunBot) ([]apiv1.ConfigMap, error) {
	list, err := r.cli.Core().ConfigMaps(r.Namespace).List(metav1.ListOptions{
		LabelSelector: labelIDKey + "=" + r.labelIDValue(bot),
	})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func (r *Runner) getDeployment(bot *model.RunBot) (*appsv1.Deployment, error) {
	list, err := r.cli.Apps().Deployments(r.Namespace).List(metav1.ListOptions{
		LabelSelector: labelIDKey + "=" + r.labelIDValue(bot),
	})
	if err != nil {
		return nil, err
	}
	if len(list.Items) == 0 {
		return nil, ErrNotExists
	}
	if len(list.Items) > 1 {
		return nil, ErrMultipleDeployments
	}
	return &list.Items[0], nil
}

func (r *Runner) makeConfigs(bot *model.RunBot) []apiv1.ConfigMap {
	confs := []apiv1.ConfigMap{}
	for _, conf := range bot.Configs {
		confs = append(confs, apiv1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name: r.kubeConfigName(bot, conf),
				Labels: map[string]string{
					labelIDKey:     r.labelIDValue(bot),
					labelConfigKey: conf.Name,
				},
			},
			Data: map[string]string{
				conf.File: conf.Value,
			},
		})
	}
	return confs
}

func (r *Runner) makeDeployment(bot *model.RunBot) *appsv1.Deployment {
	envs := []apiv1.EnvVar{}
	for _, env := range bot.Envs {
		envs = append(envs, apiv1.EnvVar{
			Name:  env.Name,
			Value: env.Value,
		})
	}

	vols := []apiv1.Volume{}
	mounts := []apiv1.VolumeMount{}
	for _, conf := range bot.Configs {
		vols = append(vols, apiv1.Volume{
			Name: r.kubeConfigName(bot, conf),
			VolumeSource: apiv1.VolumeSource{
				ConfigMap: &apiv1.ConfigMapVolumeSource{
					LocalObjectReference: apiv1.LocalObjectReference{
						Name: r.kubeConfigName(bot, conf),
					},
				},
			},
		})
		mounts = append(mounts, apiv1.VolumeMount{
			Name:      r.kubeConfigName(bot, conf),
			MountPath: filepath.Join(conf.Path, conf.File),
			SubPath:   conf.File,
		})
	}

	for _, vol := range bot.Volumes {
		vols = append(vols, apiv1.Volume{
			Name: r.kubeVolumeName(bot, vol),
			VolumeSource: apiv1.VolumeSource{
				NFS: &apiv1.NFSVolumeSource{
					Server: r.volumeManager.NfsAddr,
					Path:   r.volumeManager.GetPath(vol),
				},
			},
		})
		mounts = append(mounts, apiv1.VolumeMount{
			Name:      r.kubeVolumeName(bot, vol),
			MountPath: vol.Path,
		})
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      r.kubeDeploymentName(bot),
			Namespace: r.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					labelIDKey: r.labelIDValue(bot),
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						labelIDKey: r.labelIDValue(bot),
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:         "bot",
							Image:        bot.BuildResult,
							VolumeMounts: mounts,
							Env:          envs,
						},
					},
					Volumes: vols,
				},
			},
		},
	}
}

func (r *Runner) kubeBotName(bot *model.RunBot) string {
	return "fws-" + bot.Name + strconv.Itoa(bot.ID)
}

func (r *Runner) kubeDeploymentName(bot *model.RunBot) string {
	return r.kubeBotName(bot) + "-deployment"
}

func (r *Runner) kubeConfigName(bot *model.RunBot, conf *model.BotConfig) string {
	return r.kubeBotName(bot) + "-config-" + conf.Name
}

func (r *Runner) kubeVolumeName(bot *model.RunBot, vol *model.BotVolume) string {
	return r.kubeBotName(bot) + "-volume-" + vol.Name
}

func (r *Runner) labelIDValue(bot *model.RunBot) string {
	return strconv.Itoa(bot.ID)
}
