package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type Heal1 struct {
	iconImage int
}

func (h *Heal1) Init() {
	h.iconImage = dxlib.LoadGraph("data/images/heal.png")
	if h.iconImage == -1 {
		system.FailWithError("Failed to load heal image")
	}
}

func (h *Heal1) End() {
	dxlib.DeleteGraph(h.iconImage)
}

func (h *Heal1) Exec(manager models.Manager) {
	// WIP: HPを回復する
}

func (h *Heal1) GetParam() models.PlayerSkillParam {
	return models.PlayerSkillParam{
		CastTime:   120,
		RecastTime: 150,
		Power:      30,
		Range:      -1,
	}
}

func (h *Heal1) GetIcon() int {
	return h.iconImage
}
