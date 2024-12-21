package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

const (
	attackStateCast = iota
	attackStateSign
)

const (
	attackCastTime = 60
)

type Attack struct {
	state   int
	count   int
	ownerID string

	manager models.Manager
}

func (a *Attack) Init(manager models.Manager, ownerID string) {
	a.state = attackStateCast
	a.manager = manager
	a.ownerID = ownerID
}

func (a *Attack) End() {
}

func (a *Attack) Draw() {
	switch a.state {
	case attackStateCast:
		// 詠唱バー
		castTime := attackCastTime - a.count
		if castTime > 0 {
			size := 200
			posList := a.manager.GetPosList(&models.ObjectFilter{ID: a.ownerID})
			if len(posList) == 0 {
				return
			}
			px := posList[0].X - size/2
			py := posList[0].Y + 50
			dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
			castSize := size * castTime / attackCastTime
			dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
			dxlib.DrawFormatString(px, py+25, 0xffffff, "範囲攻撃")
		}
	case attackStateSign:
	}
}

func (a *Attack) Update() bool {
	switch a.state {
	case attackStateCast:
		// 詠唱
		if a.count >= attackCastTime {
			a.count = 0
			a.state = attackStateSign
			return false
		}
	case attackStateSign:
		// 攻撃
		return true
	}

	a.count++
	return false
}
