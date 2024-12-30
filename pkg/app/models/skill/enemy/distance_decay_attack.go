package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type DistanceDecayAttack struct {
	CastTime int
	Name     string
	FixedPos *point.Point

	count     int
	ownerID   string
	manager   models.Manager
	attackPos point.Point
}

func (a *DistanceDecayAttack) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *DistanceDecayAttack) End() {
}

func (a *DistanceDecayAttack) Draw() {
	color := dxlib.GetColor(245, 130, 32)

	// 中央のマーク
	dxlib.DrawCircle(a.attackPos.X, a.attackPos.Y, 30, color, false)
	dxlib.DrawCircle(a.attackPos.X, a.attackPos.Y, 15, color, false)

	// ライン
	minR := 40
	spd := 8
	dxlib.DrawCircle(a.attackPos.X, a.attackPos.Y, minR+spd*(a.count%80), color, false)
}

func (a *DistanceDecayAttack) Update() bool {
	if a.count == 0 {
		if a.FixedPos != nil {
			a.attackPos = *a.FixedPos
		} else {
			objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
			if len(objs) == 0 {
				return true
			}
			a.attackPos = objs[0].Pos
		}
	}

	a.count++
	return false
}

func (a *DistanceDecayAttack) GetCount() int {
	return a.count
}

func (a *DistanceDecayAttack) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: a.CastTime,
		Name:     a.Name,
	}
}
