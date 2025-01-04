package title

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/sound"
)

type Title struct {
	image int
}

func NewTitle() *Title {
	return &Title{
		image: dxlib.LoadGraph("data/images/title.png"),
	}
}

func (a *Title) Draw() {
	dxlib.DrawRotaGraph(config.ScreenSizeX/2, 180, 1, 0, a.image, true)
	dxlib.DrawFormatString(config.ScreenSizeX/2-100, 450, dxlib.GetColor(255, 255, 255), "スペースキーでスタート！")
}

func (a *Title) Update() bool {
	if dxlib.CheckHitKey(dxlib.KEY_INPUT_SPACE) == 1 {
		sound.On(sound.SEEnter)
		return true
	}
	return false
}
