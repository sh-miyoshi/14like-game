package sound

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

const (
	SEFailed int = iota

	SEMax
)

var (
	soundEffects = [SEMax]int{}
)

func Init() {
	soundEffects[SEFailed] = dxlib.LoadSoundMem("data/sounds/failed.mp3")

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
		return
	}
	dxlib.PlaySoundMem(soundEffects[typ], dxlib.DX_PLAYTYPE_BACK, true)
}
