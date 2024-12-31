package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	aeroCastTime    = 120
	aeroCenterRange = 50
)

type Aero struct {
	count     int
	ownerID   string
	manager   models.Manager
	centerPos point.Point
}

func (a *Aero) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.centerPos = point.Point{X: config.ScreenSizeX / 2, Y: config.ScreenSizeY / 2}
}

func (a *Aero) End() {
}

func (a *Aero) Draw() {
	dxlib.DrawCircle(a.centerPos.X, a.centerPos.Y, 10, dxlib.GetColor(0, 255, 0), true)
	if a.count > aeroCastTime-40 {
		dxlib.DrawCircle(a.centerPos.X, a.centerPos.Y, aeroCenterRange, dxlib.GetColor(0, 0, 255), true)
	}
}

func (a *Aero) Update() bool {
	a.count++

	if a.count == aeroCastTime {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
		for _, obj := range objs {
			d2 := point.Distance2(a.centerPos, obj.Pos)
			r := aeroCenterRange + config.PlayerHitRange
			if d2 <= r*r {
				// 中心はアウト
				a.manager.AddDamage(models.Damage{
					ID:         uuid.New().String(),
					Power:      1,
					DamageType: models.DamageTypeObject,
					TargetID:   obj.ID,
				})
			} else {
				// それ以外なら吹き飛ばし
				a.manager.AddDamage(models.Damage{
					ID:         uuid.New().String(),
					Power:      0,
					DamageType: models.DamageTypeObject,
					Push: &models.DamagePush{
						At:     a.centerPos,
						Length: 120,
						IsBack: true,
					},
					TargetID: obj.ID,
				})
			}
		}
		return true
	}
	return false
}

func (a *Aero) GetCount() int {
	return a.count
}

func (a *Aero) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: aeroCastTime,
		Name:     "エアロジャ",
	}
}
