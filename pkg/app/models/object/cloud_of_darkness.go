package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/enemy"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type CloudOfDarkness struct {
	id       string
	pos      point.Point
	timeline []struct {
		triggerTime int
		info        models.EnemySkill
	}
	count        int
	currentSkill models.EnemySkill
	manager      models.Manager
}

func (e *CloudOfDarkness) Init(manager models.Manager) {
	e.id = uuid.New().String()
	e.pos.X = config.ScreenSizeX / 2
	e.pos.Y = 50
	e.manager = manager
	e.currentSkill = nil

	e.timeline = []struct {
		triggerTime int
		info        models.EnemySkill
	}{
		{1, &skill.WaveGun{}},
	}
}

func (e *CloudOfDarkness) End() {
}

func (e *CloudOfDarkness) Draw() {
	if e.currentSkill != nil {
		e.currentSkill.Draw()
	}

	dxlib.DrawCircle(e.pos.X, e.pos.Y, 100, dxlib.GetColor(255, 255, 255), false)
	e.drawCastBar()
}

func (e *CloudOfDarkness) Update() bool {
	if e.currentSkill != nil {
		if e.currentSkill.Update() {
			e.currentSkill.End()
			e.currentSkill = nil
		}
		return false
	}

	e.count++
	for i, s := range e.timeline {
		if s.triggerTime*60 == e.count {
			logger.Debug("CloudOfDarkness trigger skill %d", i)
			e.currentSkill = s.info
			e.currentSkill.Init(e.manager, e.id)
			break
		}
	}
	return false
}

func (e *CloudOfDarkness) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       e.id,
		Pos:      e.pos,
		IsPlayer: false,
	}
}

func (e *CloudOfDarkness) HandleDamage(dm models.Damage) {
	logger.Debug("CloudOfDarkness got damage %d", dm.Power)
}

func (e *CloudOfDarkness) drawCastBar() {
	if e.currentSkill == nil {
		return
	}

	pm := e.currentSkill.GetParam()
	castTime := pm.CastTime - e.currentSkill.GetCount()
	if e.currentSkill.GetCount() != 0 && castTime > 0 {
		size := 200
		px := config.ScreenSizeX/2 - size/2
		py := config.ScreenSizeY - 100
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / pm.CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, pm.Name)
	}
}
