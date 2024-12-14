package object

import (
	"fmt"

	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/skill"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/inputs"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	PlayerHitRange = 50
	PlayerSkillMax = 3
)

type Player struct {
	pos         point.Point
	skills      [PlayerSkillMax]skill.Skill
	targetEnemy Object
	// hp          int
}

func (p *Player) Init() {
	p.pos.X = config.ScreenSizeX / 4
	p.pos.Y = config.ScreenSizeY / 2
	p.skills[0] = &skill.Attack1{}
	p.skills[1] = &skill.Heal1{}
	p.skills[2] = &skill.Defense1{}

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

func (p *Player) SetTargetEnemy(e Object) {
	p.targetEnemy = e
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
			// 使えない場合はグレーにする
			if !p.availableByDistance(s) {
				dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 160)
				dxlib.DrawBox(x, y, x+size, y+size, dxlib.GetColor(0, 0, 0), true)
				dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
			}
		}
		dxlib.DrawStringToHandle(x, y, 0xffffff, config.SkillNumberFontHandle, fmt.Sprintf("%d", i+1))
	}
}

func (p *Player) Update() {
	// WIP: スキル発動

	// Move
	spd := 4

	moveLR := 0
	moveUD := 0
	// WIP: 移動ガード
	if inputs.CheckKey(inputs.KeyUp) > 0 {
		moveUD = -spd
	} else if inputs.CheckKey(inputs.KeyDown) > 0 {
		moveUD = spd
	}

	if inputs.CheckKey(inputs.KeyRight) > 0 {
		moveLR = spd
	} else if inputs.CheckKey(inputs.KeyLeft) > 0 {
		moveLR = -spd
	}
	if moveLR != 0 && moveUD != 0 {
		// NOTE: 本来は√2で割るべきだが、見栄え的な観点で1.2にしている
		moveLR = int(float64(moveLR) / 1.2)
		moveUD = int(float64(moveUD) / 1.2)
	}
	p.pos.X += moveLR
	p.pos.Y += moveUD
}

func (p *Player) GetPos() point.Point {
	return p.pos
}

func (p *Player) availableByDistance(s skill.Skill) bool {
	if s.GetParam().Range < 0 {
		return true
	}

	if p.targetEnemy == nil {
		return false
	}

	px := p.pos.X
	py := p.pos.Y
	ex := p.targetEnemy.GetPos().X
	ey := p.targetEnemy.GetPos().Y

	dist2 := (px-ex)*(px-ex) + (py-ey)*(py-ey)
	hitRange := PlayerHitRange + s.GetParam().Range + Enemy1HitRange

	return dist2 < hitRange*hitRange
}
