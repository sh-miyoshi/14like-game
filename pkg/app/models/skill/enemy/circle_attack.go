package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type CircleAttack struct {
	CastTime int
	Range    int
	Name     string

	count     int
	ownerID   string
	attackPos point.Point
	manager   models.Manager
}

func (a *CircleAttack) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *CircleAttack) End() {
}

func (a *CircleAttack) Draw() {
	// 範囲
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	dxlib.DrawCircle(a.attackPos.X, a.attackPos.Y, a.Range, dxlib.GetColor(255, 255, 0), true)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *CircleAttack) Update() bool {
	if a.count == 0 {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
		if len(objs) == 0 {
			return true
		}
		a.attackPos = objs[0].Pos
	}

	// 詠唱
	if a.count >= a.CastTime {
		a.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      10,
			DamageType: models.DamageTypeAreaCircle,
			CenterPos:  a.attackPos,
			Range:      a.Range,
		})
		return true
	}

	a.count++
	return false
}

func (a *CircleAttack) GetCount() int {
	return a.count
}

func (a *CircleAttack) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: a.CastTime,
		Name:     a.Name,
	}
}
