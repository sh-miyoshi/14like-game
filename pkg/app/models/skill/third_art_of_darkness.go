package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

const (
	thirdArtOfDarknessCastTime = 120
)

type ThirdArtOfDarkness struct {
	count int
}

func (p *ThirdArtOfDarkness) Init(manager models.Manager, ownerID string) {
	p.count = 0
}

func (p *ThirdArtOfDarkness) End() {}

func (p *ThirdArtOfDarkness) Draw() {}

func (p *ThirdArtOfDarkness) Update() bool {
	p.count++
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
