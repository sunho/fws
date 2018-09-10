package main

import (
	"io/ioutil"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Addr       string `yaml:"addr"`
	Dist       string `yaml:"dist"`
	Dev        bool   `yaml:"dev"`
	Secret     string `yaml:"secret"`
	SqliteFile string `yaml:"sqlite_file"`
	RegURL     string `yaml:"reg_url"`
	Workspace  string `yaml:"workspace"`
}

func loadConfig() (Config, error) {
	_, err := os.Stat("config.yaml")
	if os.IsNotExist(err) {
		err = createConfig()
		if err != nil {
			return Config{}, err
		}
	} else if err != nil {
		return Config{}, err
	}

	var conf Config
	buf, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return Config{}, err
	}

	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}

func createConfig() error {
	conf := Config{
		Addr:       ":8080",
		Dist:       "dist",
		Dev:        false,
		Secret:     "thisshouldbe16lt",
		SqliteFile: "fws.db",
		Workspace:  "./workspace",
	}

	buf, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config.yaml", buf, 0644)
	if err != nil {
		return err
	}

	return nil
}
