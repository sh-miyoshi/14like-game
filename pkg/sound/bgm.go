package sound

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
)

var (
	bgmHandle int = -1
)

func BGMPlay() {
	if !config.Get().Sound.BGMEnabled {
		return
	}

	logger.Info("Start BGM")
	bgmHandle = dxlib.LoadSoundMem("data/sounds/bgm/bgm.mp3")
	if bgmHandle == -1 {
		system.FailWithError("Failed to load bgm")
	}
	dxlib.ChangeVolumeSoundMem(192, bgmHandle)

	dxlib.PlaySoundMem(bgmHandle, dxlib.DX_PLAYTYPE_LOOP, true)
}

func BGMStop() {
	if !config.Get().Sound.BGMEnabled {
		return
	}

	logger.Info("Stop BGM")
	if bgmHandle != -1 && dxlib.CheckSoundMem(bgmHandle) == 1 {
		dxlib.StopSoundMem(bgmHandle)
	}
}
