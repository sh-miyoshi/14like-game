package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	attackCastTime = 160
	attackRange    = 110
)

type Attack struct {
	count     int
	ownerID   string
	attackPos point.Point
	manager   models.Manager
}

func (a *Attack) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *Attack) End() {
}

func (a *Attack) Draw() {
	// 詠唱バー
	castTime := attackCastTime - a.count
	if a.count != 0 && castTime > 0 {
		size := 200
		objs := a.manager.GetObjectParams(&models.ObjectFilter{ID: a.ownerID})
		if len(objs) == 0 {
			return
		}
		px := objs[0].Pos.X - size/2
		py := objs[0].Pos.Y + 50
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / attackCastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "範囲攻撃")
	}

	// 範囲
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	dxlib.DrawCircle(a.attackPos.X, a.attackPos.Y, attackRange, dxlib.GetColor(255, 255, 0), true)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *Attack) Update() bool {
	if a.count == 0 {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
		if len(objs) == 0 {
			return true
		}
		a.attackPos = objs[0].Pos
	}

	// 詠唱
	if a.count >= attackCastTime {
		a.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      10,
			DamageType: models.TypeAreaCircle,
			CenterPos:  a.attackPos,
			Range:      attackRange,
		})
		return true
	}

	a.count++
	return false
}
