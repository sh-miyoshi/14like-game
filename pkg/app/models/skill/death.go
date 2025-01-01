package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	deathCastTime  = 120
	deathCastTime2 = deathCastTime + 30
	deathCastTime3 = deathCastTime2 + 30
	deathHitRange  = 50
)

type Death struct {
	count     int
	ownerID   string
	manager   models.Manager
	centerPos point.Point
}

func (a *Death) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.centerPos = point.Point{X: config.ScreenSizeX / 2, Y: config.ScreenSizeY / 2}
}

func (a *Death) End() {
}

func (a *Death) Draw() {
	dxlib.DrawCircle(a.centerPos.X, a.centerPos.Y, 10, dxlib.GetColor(185, 122, 87), true)
	if a.count >= deathCastTime && a.count < deathCastTime2 {
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		dxlib.DrawCircle(a.centerPos.X, a.centerPos.Y, deathHitRange, dxlib.GetColor(255, 255, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	} else if a.count >= deathCastTime2 {
		// WIP
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		dxlib.DrawCircle(a.centerPos.X, a.centerPos.Y, deathHitRange, dxlib.GetColor(255, 0, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}
}

func (a *Death) Update() bool {
	a.count++
	if a.count == deathCastTime {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
		for _, obj := range objs {
			a.manager.AddDamage(models.Damage{
				ID:         uuid.New().String(),
				Power:      0,
				DamageType: models.DamageTypeObject,
				Push: &models.DamagePush{
					At:     a.centerPos,
					Length: 120,
					IsBack: false,
				},
				TargetID: obj.ID,
			})
		}
	}
	if a.count == deathCastTime2 {
		a.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      1,
			DamageType: models.DamageTypeAreaCircle,
			CenterPos:  a.centerPos,
			Range:      deathHitRange,
		})
	}

	return false
}

func (a *Death) GetCount() int {
	return a.count
}

func (a *Death) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: deathCastTime,
		Name:     "デスジャ",
	}
}
