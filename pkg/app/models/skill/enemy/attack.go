package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
)

const (
	attackStateCast = iota
	attackStateSign
)

const (
	attackCastTime = 60
)

type Attack struct {
	state int
	count int

	manager models.Manager
}

func (a *Attack) Init(manager models.Manager) {
	a.state = attackStateCast
	a.manager = manager
}

func (a *Attack) End() {
}

func (a *Attack) Draw() {
	switch a.state {
	case attackStateCast:
		// 詠唱バー
		// if e.castTime > 0 {
		// 	size := 200
		// 	px := e.pos.X - size/2
		// 	py := e.pos.Y + Enemy1HitRange + 30
		// 	dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		// 	castSize := size * e.castTime / e.currentSkill.GetParam().CastTime
		// 	dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		// 	dxlib.DrawFormatString(px, py+25, 0xffffff, e.currentSkill.GetParam().Name)
		// }
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
	}

	a.count++
	return false
}
