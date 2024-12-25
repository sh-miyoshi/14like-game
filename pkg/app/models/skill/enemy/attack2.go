package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	attack2CastTime = 260
)

type Attack2 struct {
	count     int
	ownerID   string
	attackPos [4]point.Point
	manager   models.Manager
}

func (a *Attack2) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *Attack2) End() {
}

func (a *Attack2) Draw() {
	// 詠唱バー
	castTime := attack2CastTime - a.count
	if a.count != 0 && castTime > 0 {
		size := 200
		posList := a.manager.GetPosList(&models.ObjectFilter{ID: a.ownerID})
		if len(posList) == 0 {
			return
		}
		px := posList[0].X - size/2
		py := posList[0].Y + 50
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / attack2CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "範囲攻撃2")
	}

	// 範囲
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	dxlib.DrawQuadrangle(
		a.attackPos[0].X,
		a.attackPos[0].Y,
		a.attackPos[1].X,
		a.attackPos[1].Y,
		a.attackPos[2].X,
		a.attackPos[2].Y,
		a.attackPos[3].X,
		a.attackPos[3].Y,
		dxlib.GetColor(255, 255, 0),
		true,
	)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *Attack2) Update() bool {
	if a.count == 0 {
		// WIP: とりあえず固定の位置に攻撃範囲を設定
		// WIP: Rotate
		a.attackPos[0] = point.Point{X: 100, Y: 100}
		a.attackPos[1] = point.Point{X: 100, Y: 500}
		a.attackPos[2] = point.Point{X: 600, Y: 500}
		a.attackPos[3] = point.Point{X: 600, Y: 100}
	}

	// 詠唱
	if a.count >= attack2CastTime {
		a.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      10,
			DamageType: models.TypeAreaRect,
			RectPos:    [2]point.Point{a.attackPos[0], a.attackPos[2]},
		})
		return true
	}

	a.count++
	return false
}
