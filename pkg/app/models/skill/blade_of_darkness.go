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
	bladeOfDarknessCastTime = 180
)

const (
	BladeOfDarknessAttackLeft int = iota
	BladeOfDarknessAttackRight
	BladeOfDarknessAttackCenter
)

type BladeOfDarkness struct {
	AttackType int

	count   int
	ownerID string
	manager models.Manager
	image   int
}

func (a *BladeOfDarkness) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.image = dxlib.LoadGraph("data/images/blade_of_darkness_area.png")
	if a.image == -1 {
		system.FailWithError("Failed to load blade_of_darkness_area image")
	}
}

func (a *BladeOfDarkness) End() {
	dxlib.DeleteGraph(a.image)
}

func (a *BladeOfDarkness) Draw() {
	color := dxlib.GetColor(34, 176, 84)
	switch a.AttackType {
	case BladeOfDarknessAttackLeft:
		dxlib.DrawCircle(250, 200, 30, color, true)
	case BladeOfDarknessAttackRight:
		dxlib.DrawCircle(550, 200, 30, color, true)
	case BladeOfDarknessAttackCenter:
		size := 120
		t := int32(5)
		dxlib.DrawLine(
			config.ScreenSizeX/2-size/2, 100,
			config.ScreenSizeX/2+size/2, 100+size,
			color,
			dxlib.DrawLineOption{
				Thickness: &t,
			},
		)
		dxlib.DrawLine(
			config.ScreenSizeX/2-size/2, 100+size,
			config.ScreenSizeX/2+size/2, 100,
			color,
			dxlib.DrawLineOption{
				Thickness: &t,
			},
		)
	}

	// WIP: タイミング
	dxlib.SetDrawArea(0, 200, config.ScreenSizeX, config.ScreenSizeY)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	switch a.AttackType {
	case BladeOfDarknessAttackLeft:
		dxlib.DrawRotaGraph(250, 200, 2, 0, a.image, true)
	case BladeOfDarknessAttackRight:
		dxlib.DrawRotaGraph(550, 200, 2, 0, a.image, true)
	case BladeOfDarknessAttackCenter:
		dxlib.DrawCircle(config.ScreenSizeX/2, 200, 250, dxlib.GetColor(255, 255, 0), true)
	}
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	dxlib.SetDrawArea(0, 0, config.ScreenSizeX, config.ScreenSizeY)
}

func (a *BladeOfDarkness) Update() bool {
	a.count++
	if a.count == bladeOfDarknessCastTime {
		switch a.AttackType {
		case BladeOfDarknessAttackLeft:
			a.manager.AddDamage(models.Damage{
				ID:         uuid.New().String(),
				Power:      1,
				DamageType: models.DamageTypeAreaRing,
				CenterPos:  point.Point{X: 250, Y: 200},
				Range:      500,
				InnerRange: 90,
			})
		case BladeOfDarknessAttackRight:
			a.manager.AddDamage(models.Damage{
				ID:         uuid.New().String(),
				Power:      1,
				DamageType: models.DamageTypeAreaRing,
				CenterPos:  point.Point{X: 550, Y: 200},
				Range:      500,
				InnerRange: 90,
			})
		case BladeOfDarknessAttackCenter:
			a.manager.AddDamage(models.Damage{
				ID:         uuid.New().String(),
				Power:      1,
				DamageType: models.DamageTypeAreaCircle,
				CenterPos:  point.Point{X: config.ScreenSizeX / 2, Y: 200},
				Range:      250,
			})
		}
		return true
	}
	return false
}

func (a *BladeOfDarkness) GetCount() int {
	return a.count
}

func (a *BladeOfDarkness) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: bladeOfDarknessCastTime,
		Name:     "闇の刃",
	}
}
