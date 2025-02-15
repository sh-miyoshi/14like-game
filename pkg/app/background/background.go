package background

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type BackGround struct {
	phase int
}

func (b *BackGround) Init(phase int) {
	b.phase = phase
}

func (b *BackGround) Draw() {
	switch b.phase {
	case config.Phase1:
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
	case config.Phase2B:
		size := 60
		ofsx := config.ScreenSizeX/2 - 4*size
		ofsy := 60
		for x := 0; x < 8; x++ {
			y := ofsy + size
			dxlib.DrawBox(x*size+ofsx, y, x*size+size+ofsx, y+size, 0xFFFFFF, false)
			y = ofsy + 6*size
			dxlib.DrawBox(x*size+ofsx, y, x*size+size+ofsx, y+size, 0xFFFFFF, false)
		}
		for y := 0; y < 8; y++ {
			x := ofsx + size
			dxlib.DrawBox(x, y*size+ofsy, x+size, y*size+size+ofsy, 0xFFFFFF, false)
			x = ofsx + 6*size
			dxlib.DrawBox(x, y*size+ofsy, x+size, y*size+size+ofsy, 0xFFFFFF, false)
		}
	}
}

func (b *BackGround) Update() {
}
