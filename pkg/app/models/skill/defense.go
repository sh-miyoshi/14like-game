package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type Defense1 struct {
	iconImage int
}

func (h *Defense1) Init() {
	h.iconImage = dxlib.LoadGraph("data/images/defense.png")
	if h.iconImage == -1 {
		system.FailWithError("Failed to load defense image")
	}
}

func (h *Defense1) End() {
	dxlib.DeleteGraph(h.iconImage)
}

func (h *Defense1) GetParam() Param {
	return Param{
		CastTime:   0,
		RecastTime: 180,
		Power:      30,
		Range:      -1,
	}
}

func (h *Defense1) GetIcon() int {
	return h.iconImage
}
