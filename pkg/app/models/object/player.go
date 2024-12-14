package object

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/skill"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	PlayerHitRange = 50
	PlayerSkillMax = 4
)

type Player struct {
	pos    point.Point
	skills [PlayerSkillMax]skill.Skill
}

func (p *Player) Init() {
	p.pos.X = config.ScreenSizeY / 4
	p.pos.Y = config.ScreenSizeY / 2
	p.skills[0] = &skill.Attack1{}

	for _, s := range p.skills {
		if s != nil {
			s.Init()
		}
	}
}

func (p *Player) End() {
	for _, s := range p.skills {
		if s != nil {
			s.End()
		}
	}
}

func (p *Player) Draw() {
	dxlib.DrawCircle(p.pos.X, p.pos.Y, PlayerHitRange, dxlib.GetColor(255, 255, 255), false)

	for i, s := range p.skills {
		size := 32
		x := i*(size+15) + 35
		y := config.ScreenSizeY - 60
		if s == nil {
			dxlib.DrawBox(x, y, x+size, y+size, dxlib.GetColor(255, 255, 255), false)
		} else {
			dxlib.DrawGraph(x, y, s.GetIcon(), true)
			// WIP: 使えない場合はグレーにする
		}
	}
}

func (p *Player) Update() {
}
