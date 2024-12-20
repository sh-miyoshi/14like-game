package object

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/enemy"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	EnemyHPSize    = 200
	Enemy1HitRange = 150
)

type enemySkill struct {
	info        skill.Skill
	triggerTime int
}

type Enemy1 struct {
	pos          point.Point
	hp           int
	hpMax        int
	timeline     []enemySkill
	count        int
	castTime     int
	currentSkill skill.Skill
	addDamage    func(models.Damage)
}

func (e *Enemy1) Init(addDamage func(models.Damage)) {
	e.pos.X = config.ScreenSizeX * 3 / 4
	e.pos.Y = config.ScreenSizeY / 2
	e.hpMax = 1000
	e.hp = e.hpMax
	e.timeline = []enemySkill{
		{triggerTime: 100, info: &skill.Attack{}},
	}
	e.addDamage = addDamage
}

func (e *Enemy1) End() {
}

func (e *Enemy1) Draw() {
	dxlib.DrawCircle(e.pos.X, e.pos.Y, Enemy1HitRange, dxlib.GetColor(255, 255, 255), false)

	tx := config.ScreenSizeX - EnemyHPSize - 40
	ty := 30
	dxlib.DrawBox(tx, ty, tx+EnemyHPSize, ty+20, dxlib.GetColor(255, 255, 255), false)
	size := EnemyHPSize * e.hp / e.hpMax
	dxlib.DrawBox(tx, ty, tx+size, ty+20, dxlib.GetColor(255, 255, 255), true)

	// 詠唱バー
	if e.castTime > 0 {
		size := 200
		px := e.pos.X - size/2
		py := e.pos.Y + Enemy1HitRange + 30
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * e.castTime / e.currentSkill.GetParam().CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, e.currentSkill.GetParam().Name)
	}
}

func (e *Enemy1) Update() {
	if e.castTime > 0 {
		e.castTime--
		if e.castTime == 0 {
			e.currentSkill.Exec(e.addDamage)
		}
	}

	e.count++
	for i, s := range e.timeline {
		if s.triggerTime == e.count {
			logger.Debug("Enemy1 trigger skill %d", i)
			e.castTime = s.info.GetParam().CastTime
			e.currentSkill = s.info
			// WIP: castTimeが0の技
			break
		}
	}
}

func (e *Enemy1) GetPos() point.Point {
	return e.pos
}

func (e *Enemy1) IsPlayer() bool {
	return false
}

func (e *Enemy1) HandleDamage(power int) {
	logger.Debug("Enemy1 got damage %d", power)
	e.hp -= power
	if e.hp < 0 {
		e.hp = 0
	}
}
