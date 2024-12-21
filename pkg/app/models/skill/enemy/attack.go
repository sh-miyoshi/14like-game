package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
)

const (
	attackStateCast = iota
	attackStateSign
)

type Attack struct {
	state int
	count int
}

func (a *Attack) Init() {
	a.state = attackStateCast
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

func (a *Attack) Update(manager models.Manager) bool {
	switch a.state {
	case attackStateCast:
		// 詠唱
		a.state = attackStateSign
	case attackStateSign:
		// 攻撃
	}

	a.count++
	return false
}
