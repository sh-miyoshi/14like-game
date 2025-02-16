package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type OnlyCast struct {
	CastTime int
	Name     string
	Text     string

	count   int
	ownerID string
	manager models.Manager
}

func (a *OnlyCast) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *OnlyCast) End() {
}

func (a *OnlyCast) Draw() {
	if a.Text != "" {
		ofs := dxlib.GetDrawStringWidth(a.Text, len(a.Text))
		dxlib.DrawFormatString(config.ScreenSizeX/2-ofs/2, config.ScreenSizeY/2-20, dxlib.GetColor(255, 255, 255), "%s", a.Text)
	}
}

func (a *OnlyCast) Update() bool {
	a.count++
	return a.count >= a.CastTime
}

func (a *OnlyCast) GetCount() int {
	return a.count
}

func (a *OnlyCast) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: a.CastTime,
		Name:     a.Name,
	}
}
