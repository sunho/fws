package basic

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type VolumeManager struct {
	NfsPath string
	NfsAddr string
}

func (v *VolumeManager) getBotTashDir(bot int) string {
	return filepath.Join("/trash", strconv.Itoa(bot))
}

func (v *VolumeManager) getBotDir(bot int) string {
	return filepath.Join("/", strconv.Itoa(bot))
}

func (v *VolumeManager) actualize(path string) string {
	return filepath.Join(v.NfsPath, path)
}

func (v *VolumeManager) GetTrashPath(bot int, vol string) string {
	return filepath.Join(v.getBotTashDir(bot), vol+"-"+strconv.Itoa(int(time.Now().Unix())))
}

func (v *VolumeManager) GetPath(bot int, vol string) string {
	return filepath.Join(v.getBotDir(bot), vol)
}

func (v *VolumeManager) initDir(path string) error {
	info, err := os.Stat(v.actualize(path))
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
		return nil
	} else if err != nil {
		return err
	}

	if !info.IsDir() {
		err := os.Remove(path)
		if err != nil {
			return err
		}

		err = os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v *VolumeManager) Init(bot int) error {
	err := v.initDir(v.getBotDir(bot))
	if err != nil {
		return err
	}
	err = v.initDir(v.getBotTashDir(bot))
	if err != nil {
		return err
	}
	return nil
}

func (v *VolumeManager) List(bot int) ([]string, error) {
	files, err := ioutil.ReadDir(v.actualize(v.getBotDir(bot)))
	if err != nil {
		return nil, err
	}
	out := make([]string, 0, len(files))
	for _, f := range files {
		out = append(out, f.Name())
	}

	return out, nil
}

func (v *VolumeManager) Stats(bot int, vol string) bool {
	_, err := os.Stat(v.actualize(v.GetPath(bot, vol)))
	return !os.IsNotExist(err)
}

func (v *VolumeManager) Create(bot int, vol string) error {
	return os.MkdirAll(v.actualize(v.GetPath(bot, vol)), 0777)
}

func (v *VolumeManager) Delete(bot int, vol string) error {
	return os.Rename(v.actualize(v.GetPath(bot, vol)), v.actualize(v.GetTrashPath(bot, vol)))
}
