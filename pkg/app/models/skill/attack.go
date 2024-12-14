package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type Attack1 struct {
	iconImage int
}

func (a *Attack1) Init() {
	a.iconImage = dxlib.LoadGraph("data/images/attack.png")
	if a.iconImage == -1 {
		system.FailWithError("Failed to load attack image")
	}
}

func (a *Attack1) End() {
	dxlib.DeleteGraph(a.iconImage)
}

func (a *Attack1) GetParam() Param {
	// WIP
	return Param{}
}

func (a *Attack1) GetIcon() int {
	return a.iconImage
}
