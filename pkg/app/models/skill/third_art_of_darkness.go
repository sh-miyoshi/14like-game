package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	thirdArtOfDarknessCastTime     = 30
	thirdArtOfDarknessCastInterval = 60
)

const (
	thirdArtOfDarknessAttackTypeLeft int = iota
	thirdArtOfDarknessAttackTypeRight
	thirdArtOfDarknessAttackTypeSpread
	thirdArtOfDarknessAttackTypeGroup
)

type ThirdArtOfDarkness struct {
	count    int
	attacks  [3]int
	ownerPos point.Point
}

func (p *ThirdArtOfDarkness) Init(manager models.Manager, ownerID string) {
	p.count = 0
	p.ownerPos = manager.GetObjectParams(&models.ObjectFilter{ID: ownerID})[0].Pos

	// Note: Left, Rightの中から2回 + Spread, Groupの中から1回、順番はランダム
	// lr := []int{thirdArtOfDarknessAttackTypeLeft, thirdArtOfDarknessAttackTypeRight}
	// sg := []int{thirdArtOfDarknessAttackTypeSpread, thirdArtOfDarknessAttackTypeGroup}
	// p.attacks[0] = lr[rand.Intn(2)]
	// p.attacks[1] = lr[rand.Intn(2)]
	// p.attacks[2] = sg[rand.Intn(2)]
	// math.Shuffle(p.attacks[:])

	// debug
	p.attacks[0] = thirdArtOfDarknessAttackTypeLeft
	p.attacks[1] = thirdArtOfDarknessAttackTypeGroup
	p.attacks[2] = thirdArtOfDarknessAttackTypeSpread
}

func (p *ThirdArtOfDarkness) End() {
}

func (p *ThirdArtOfDarkness) Draw() {
	cnt := p.count - thirdArtOfDarknessCastTime

	// 技見せフェーズ
	if cnt >= 0 && cnt < thirdArtOfDarknessCastInterval*3 && cnt%thirdArtOfDarknessCastInterval < thirdArtOfDarknessCastInterval-10 {
		switch p.attacks[cnt/thirdArtOfDarknessCastInterval] {
		case thirdArtOfDarknessAttackTypeLeft:
			dxlib.DrawCircle(p.ownerPos.X-50, p.ownerPos.Y, 30, dxlib.GetColor(168, 88, 168), true)
		case thirdArtOfDarknessAttackTypeRight:
			dxlib.DrawCircle(p.ownerPos.X+50, p.ownerPos.Y, 30, dxlib.GetColor(168, 88, 168), true)
		case thirdArtOfDarknessAttackTypeSpread:
			dxlib.DrawCircle(p.ownerPos.X, p.ownerPos.Y-20, 30, dxlib.GetColor(168, 88, 168), true)
		case thirdArtOfDarknessAttackTypeGroup:
			dxlib.DrawCircle(p.ownerPos.X-50, p.ownerPos.Y, 30, dxlib.GetColor(168, 88, 168), true)
			dxlib.DrawCircle(p.ownerPos.X+50, p.ownerPos.Y, 30, dxlib.GetColor(168, 88, 168), true)
		}
	}
}

func (p *ThirdArtOfDarkness) Update() bool {
	p.count++
	/*
		技見せフェーズ x3
		ちょっと開く
		攻撃
	*/
	return false
}

func (p *ThirdArtOfDarkness) GetCount() int {
	return p.count
}

func (p *ThirdArtOfDarkness) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: thirdArtOfDarknessCastTime,
		Name:     "闇の戦技:三重",
	}
}
