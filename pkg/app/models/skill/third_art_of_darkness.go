package skill

import (
	"math/rand"

	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
)

const (
	thirdArtOfDarknessCastTime = 120
)

const (
	thirdArtOfDarknessAttackTypeLeft int = iota
	thirdArtOfDarknessAttackTypeRight
	thirdArtOfDarknessAttackTypeSpread
	thirdArtOfDarknessAttackTypeGroup
)

type ThirdArtOfDarkness struct {
	count   int
	attacks [3]int
}

func (p *ThirdArtOfDarkness) Init(manager models.Manager, ownerID string) {
	p.count = 0
	// Note: Left, Rightの中から2回 + Spread, Groupの中から1回、順番はランダム
	lr := []int{thirdArtOfDarknessAttackTypeLeft, thirdArtOfDarknessAttackTypeRight}
	sg := []int{thirdArtOfDarknessAttackTypeSpread, thirdArtOfDarknessAttackTypeGroup}
	p.attacks[0] = lr[rand.Intn(2)]
	p.attacks[1] = lr[rand.Intn(2)]
	p.attacks[2] = sg[rand.Intn(2)]
	math.Shuffle(p.attacks[:])
}

func (p *ThirdArtOfDarkness) End() {
}

func (p *ThirdArtOfDarkness) Draw() {
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
