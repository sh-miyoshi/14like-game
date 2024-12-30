package background

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type BackGround struct {
}

func (b *BackGround) Draw() {
	dxlib.DrawQuadrangle(
		config.ScreenSizeX/2, 50,
		150, config.ScreenSizeY/2,
		config.ScreenSizeX/2, config.ScreenSizeY-50,
		config.ScreenSizeX-150, config.ScreenSizeY/2,
		0xFFFFFF,
		false,
	)
	dxlib.DrawBox(0, 0, config.ScreenSizeX, 200, 0x000000, true)
	dxlib.DrawLine(0, 200, config.ScreenSizeX, 200, 0xFFFFFF)
}

func (b *BackGround) Update() {

}
