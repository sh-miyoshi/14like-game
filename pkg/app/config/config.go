package config

import (
	"fmt"
	"os"

	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Sound struct {
		BGMEnabled bool `yaml:"bgm_enabled"`
		SEEnabled  bool `yaml:"se_enabled"`
	} `yaml:"sound"`
}

const (
	ScreenSizeX = 800 // 1280
	ScreenSizeY = 600 // 960
)

var (
	SkillNumberFontHandle int
	conf                  Config
)

const (
	PlayerHitRange = 10
)

func Init() {
	fp, err := os.Open("config.yml")
	if err != nil {
		system.FailWithError(fmt.Sprintf("Failed to open config file: %+v", err))
	}
	defer fp.Close()

	if err := yaml.NewDecoder(fp).Decode(&conf); err != nil {
		system.FailWithError(fmt.Sprintf("Failed to decode yaml: %+v", err))
	}
}

func Get() Config {
	return conf
}
