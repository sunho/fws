package basic

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sunho/fws/server/model"
)

type VolumeManager struct {
	NfsPath string
	NfsAddr string
}

func (v *VolumeManager) GetPath(vol *model.BotVolume) string {
	return filepath.Join(v.NfsPath, strconv.Itoa(vol.BotID), vol.Name)
}

func (v *VolumeManager) Stats(vol *model.BotVolume) bool {
	_, err := os.Stat(v.GetPath(vol))
	return !os.IsNotExist(err)
}

func (v *VolumeManager) Create(vol *model.BotVolume) error {
	return os.MkdirAll(v.GetPath(vol), 0644)
}

func (v *VolumeManager) Delete(vol *model.BotVolume) error {
	return os.Rename(v.GetPath(vol), v.GetPath(vol)+strconv.Itoa(int(time.Now().Unix())))
}
