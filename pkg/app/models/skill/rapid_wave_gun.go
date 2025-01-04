package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	rapidWaveGunCastTime = 180
	rapidWaveGunEndTime  = 240
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
	h := 230
	dxlib.DrawBox(x-w/2, y, x+w/2, y+h, dxlib.GetColor(126, 132, 247), true)
	// Hit Area
	if a.count >= rapidWaveGunCastTime {
		w2 := 200
		dxlib.DrawBox(x-w/2, 200, x+w/2, y, dxlib.GetColor(255, 255, 0), true)
		dxlib.DrawBox(x-w/2-w2, 200, x-w/2, y+h, dxlib.GetColor(255, 255, 0), true)
		dxlib.DrawBox(x+w/2, 200, x+w/2+w2, y+h, dxlib.GetColor(255, 255, 0), true)
	}
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *RapidWaveGun) Update() bool {
	a.count++
	if a.count == rapidWaveGunCastTime {
		dm := models.Damage{
			ID:         uuid.New().String(),
			Power:      1,
			DamageType: models.DamageTypeAreaRect,
			RectPos: [2]point.Point{
				{X: 0, Y: 0},
				{X: config.ScreenSizeX/2 - 50, Y: config.ScreenSizeY},
			},
		}

		a.manager.AddDamage(dm)
		dm.RectPos = [2]point.Point{
			{X: config.ScreenSizeX/2 + 50, Y: 0},
			{X: config.ScreenSizeX, Y: config.ScreenSizeY},
		}
		a.manager.AddDamage(dm)
		dm.RectPos = [2]point.Point{
			{X: 0, Y: 0},
			{X: config.ScreenSizeX, Y: 300},
		}
		a.manager.AddDamage(dm)
	}

	return a.count >= rapidWaveGunEndTime
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
