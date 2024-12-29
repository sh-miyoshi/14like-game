package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/enemy"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	EnemyHPSize    = 200
	Enemy1HitRange = 150
)

type enemySkill struct {
	info        models.EnemySkill
	triggerTime int
}

type Enemy1 struct {
	id           string
	pos          point.Point
	hp           int
	hpMax        int
	timeline     []enemySkill
	count        int
	currentSkill models.EnemySkill
	manager      models.Manager
	image        int
}

func (e *Enemy1) Init(manager models.Manager) {
	e.id = uuid.New().String()
	e.pos.X = config.ScreenSizeX * 3 / 4
	e.pos.Y = config.ScreenSizeY / 2
	e.hpMax = 1000
	e.hp = e.hpMax
	e.image = dxlib.LoadGraph("data/images/objects/enemy1.png")
	if e.image == -1 {
		system.FailWithError("Failed to load enemy1 image")
	}

	e.timeline = []enemySkill{
		{triggerTime: 2, info: &skill.BombBoulderMgr{}},
		// {triggerTime: 2, info: &skill.LandSlide{AttackNum: 1}},
		// {triggerTime: 4, info: &skill.FullAttack{Name: "激震", CastTime: 30, Power: 100}},
		// {triggerTime: 8, info: &skill.LandSlide{AttackNum: 3}},
	}
	e.manager = manager
	e.currentSkill = nil
}

func (e *Enemy1) End() {
	dxlib.DeleteGraph(e.image)
}

func (e *Enemy1) Draw() {
	if e.currentSkill != nil {
		e.currentSkill.Draw()
	}

	dxlib.DrawCircle(e.pos.X, e.pos.Y, Enemy1HitRange, dxlib.GetColor(255, 255, 255), false)
	dxlib.DrawRotaGraph(e.pos.X, e.pos.Y, 1, 0, e.image, true)

	tx := config.ScreenSizeX - EnemyHPSize - 40
	ty := 30
	dxlib.DrawBox(tx, ty, tx+EnemyHPSize, ty+20, dxlib.GetColor(255, 255, 255), false)
	size := EnemyHPSize * e.hp / e.hpMax
	dxlib.DrawBox(tx, ty, tx+size, ty+20, dxlib.GetColor(255, 255, 255), true)

	e.drawCastBar()
}

func (e *Enemy1) Update() {
	if e.currentSkill != nil {
		if e.currentSkill.Update() {
			e.currentSkill.End()
			e.currentSkill = nil
		}
		return
	}

	e.count++
	for i, s := range e.timeline {
		if s.triggerTime*60 == e.count {
			logger.Debug("Enemy1 trigger skill %d", i)
			e.currentSkill = s.info
			e.currentSkill.Init(e.manager, e.id)
			break
		}
	}
}

func (e *Enemy1) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       e.id,
		Pos:      e.pos,
		IsPlayer: false,
	}
}

func (e *Enemy1) HandleDamage(dm models.Damage) {
	logger.Debug("Enemy1 got damage %d", dm.Power)
	e.hp -= dm.Power
	if e.hp < 0 {
		e.hp = 0
	}
}

func (e *Enemy1) drawCastBar() {
	if e.currentSkill == nil {
		return
	}

	pm := e.currentSkill.GetParam()
	castTime := pm.CastTime - e.currentSkill.GetCount()
	if e.currentSkill.GetCount() != 0 && castTime > 0 {
		size := 200
		px := e.pos.X - size/2
		py := e.pos.Y + Enemy1HitRange - 20
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / pm.CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, pm.Name)
	}
}
