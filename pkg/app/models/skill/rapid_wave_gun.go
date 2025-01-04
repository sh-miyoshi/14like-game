package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

const (
	rapidWaveGunCastTime = 180
)

type RapidWaveGun struct {
	count   int
	ownerID string
	manager models.Manager
	image   int
}

func (a *RapidWaveGun) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.image = dxlib.LoadGraph("data/images/tank.png")
	if a.image == -1 {
		system.FailWithError("Failed to load image")
	}
}

func (a *RapidWaveGun) End() {
	dxlib.DeleteGraph(a.image)
}

func (a *RapidWaveGun) Draw() {
	// Tank
	x := config.ScreenSizeX / 2
	y := 300
	dxlib.DrawCircle(x, y, 20, dxlib.GetColor(255, 255, 255), false)
	dxlib.DrawRotaGraph(x, y, 0.4, 0.0, a.image, true)

	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	w := 100
	h := 200
	dxlib.DrawBox(x-w/2, y, x+w/2, y+h, dxlib.GetColor(0, 0, 255), true) // WIP color
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *RapidWaveGun) Update() bool {
	a.count++
	// WIP: damage, end
	return false
}

func (a *RapidWaveGun) GetCount() int {
	return a.count
}

func (a *RapidWaveGun) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: rapidWaveGunCastTime,
		Name:     "連射式波動砲",
	}
}
