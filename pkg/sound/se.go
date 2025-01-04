package sound

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

const (
	SEFailed int = iota
	SEEnter

	SEMax
)

var (
	soundEffects = [SEMax]int{}
)

func SEInit() {
	soundEffects[SEFailed] = dxlib.LoadSoundMem("data/sounds/se/failed.mp3")
	soundEffects[SEEnter] = dxlib.LoadSoundMem("data/sounds/se/enter.mp3")

	for _, se := range soundEffects {
		if se == -1 {
			system.FailWithError("Failed to load sound effect")
		}
	}
}

func On(typ int) {
	if typ < 0 || typ >= SEMax {
		return
	}

	if dxlib.CheckSoundMem(soundEffects[typ]) == 1 {
		dxlib.StopSoundMem(soundEffects[typ])
	}
	dxlib.PlaySoundMem(soundEffects[typ], dxlib.DX_PLAYTYPE_BACK, true)
}
