package object

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/player"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/inputs"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	PlayerSkillMax = 3
)

type playerSkill struct {
	info     skill.Skill
	waitTime int
}

type Player struct {
	id             string
	pos            point.Point
	skills         [PlayerSkillMax]*playerSkill
	targetEnemy    Object
	imgSkillCircle int
	castTime       int
	castSkillIndex int
	// hp          int

	manager models.Manager
}

func (p *Player) Init(manager models.Manager) {
	p.id = uuid.New().String()
	p.imgSkillCircle = dxlib.LoadGraph("data/images/skill_circle.png")
	if p.imgSkillCircle == -1 {
		system.FailWithError("Failed to load skill circle image")
	}
	p.manager = manager

	p.castTime = 0
	p.castSkillIndex = 0
	p.pos.X = config.ScreenSizeX / 4
	p.pos.Y = config.ScreenSizeY / 2
	p.skills[0] = &playerSkill{
		info:     &skill.Attack1{},
		waitTime: 0,
	}
	p.skills[1] = &playerSkill{
		info:     &skill.Heal1{},
		waitTime: 0,
	}
	p.skills[2] = &playerSkill{
		info:     &skill.Defense1{},
		waitTime: 0,
	}

	for _, s := range p.skills {
		if s != nil {
			s.info.Init()
		}
	}
}

func (p *Player) End() {
	dxlib.DeleteGraph(p.imgSkillCircle)
	for _, s := range p.skills {
		if s != nil {
			s.info.End()
		}
	}
}

func (p *Player) SetTargetEnemy(e Object) {
	p.targetEnemy = e
}

func (p *Player) Draw() {
	dxlib.DrawCircle(p.pos.X, p.pos.Y, config.PlayerHitRange, dxlib.GetColor(255, 255, 255), false)

	for i, s := range p.skills {
		size := 32
		x := i*(size+15) + 35
		y := config.ScreenSizeY - 60
		if s == nil {
			dxlib.DrawBox(x, y, x+size, y+size, dxlib.GetColor(255, 255, 255), false)
		} else {
			dxlib.DrawGraph(x, y, s.info.GetIcon(), true)
			if s.waitTime > 0 {
				dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 160)
				dxlib.DrawBox(x, y, x+size, y+size, dxlib.GetColor(0, 0, 0), true)
				dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
				tm := float64(s.waitTime) / 60.0
				dxlib.DrawFormatString(x+size/2-13, y+size/2-3, 0xffffff, "%.1f", tm)
			}
		}
		dxlib.DrawStringToHandle(x, y, 0xffffff, config.SkillNumberFontHandle, fmt.Sprintf("%d", i+1))
	}

	// 詠唱バー
	if p.castTime > 0 {
		size := 100
		px := p.pos.X - size/2
		py := p.pos.Y + config.PlayerHitRange + 30
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * p.castTime / p.skills[p.castSkillIndex].info.GetParam().CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "CAST")
	}
}

func (p *Player) Update() {
	// スキル発動
	if p.castTime == 0 {
		if inputs.CheckKey(inputs.Key1) == 1 && p.availableByDistance(p.skills[0].info) && p.skills[0].waitTime == 0 {
			p.castTime = p.skills[0].info.GetParam().CastTime + 1
			p.castSkillIndex = 0
		}
		if inputs.CheckKey(inputs.Key2) == 1 && p.availableByDistance(p.skills[1].info) && p.skills[1].waitTime == 0 {
			p.castTime = p.skills[1].info.GetParam().CastTime + 1
			p.castSkillIndex = 1
		}
		if inputs.CheckKey(inputs.Key3) == 1 && p.availableByDistance(p.skills[2].info) && p.skills[2].waitTime == 0 {
			p.castTime = p.skills[2].info.GetParam().CastTime + 1
			p.castSkillIndex = 2
		}
	}

	if p.castTime > 0 {
		p.castTime--
		if p.castTime == 0 {
			p.skills[p.castSkillIndex].waitTime = p.skills[p.castSkillIndex].info.GetParam().RecastTime
			p.skills[p.castSkillIndex].info.Exec(p.manager)
		}
	}

	for _, s := range p.skills {
		if s != nil && s.waitTime > 0 {
			s.waitTime--
		}
	}

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
	// 動いたらキャンセル
	if moveLR != 0 || moveUD != 0 {
		p.castTime = 0
	}
	p.pos.X += moveLR
	p.pos.Y += moveUD
}

func (p *Player) GetParam() Param {
	return Param{
		ID:       p.id,
		Pos:      p.pos,
		IsPlayer: true,
	}
}

func (p *Player) HandleDamage(power int) {
	logger.Debug("Player got damage %d", power)
	// WIP: ダメージ処理
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
	ex := p.targetEnemy.GetParam().Pos.X
	ey := p.targetEnemy.GetParam().Pos.Y

	dist2 := (px-ex)*(px-ex) + (py-ey)*(py-ey)
	hitRange := config.PlayerHitRange + s.GetParam().Range + Enemy1HitRange

	return dist2 < hitRange*hitRange
}
