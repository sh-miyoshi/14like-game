package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/skill"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type StygianShadow struct {
	id           string
	pos          point.Point
	timeline     []SkillTimeline
	count        int
	currentSkill models.Skill
	manager      models.Manager
	image        int
}

func (e *StygianShadow) Init(manager models.Manager) {
	e.id = uuid.New().String()
	e.manager = manager
	e.currentSkill = nil
	e.image = dxlib.LoadGraph("data/images/objects/stygian_shadow.png")
	if e.image == -1 {
		system.FailWithError("Failed to load stygian_shadow image")
	}
	e.pos = point.Point{X: config.ScreenSizeX / 2, Y: 200}

	e.timeline = []SkillTimeline{
		{60, &skill.OnlyCast{CastTime: 120, Name: "闇の大氾濫", Text: "おわり～"}},
	}
}

func (e *StygianShadow) End() {
}

func (e *StygianShadow) Draw() {
	dxlib.DrawRotaGraph(e.pos.X, e.pos.Y, 0.3, 0.0, e.image, true)
	if e.currentSkill != nil {
		e.currentSkill.Draw()
	}

	e.drawCastBar()
}

func (e *StygianShadow) Update() bool {
	if e.currentSkill != nil {
		if e.currentSkill.Update() {
			e.currentSkill.End()
			e.currentSkill = nil
		}
		return false
	}

	e.count++
	for i, s := range e.timeline {
		if s.TriggerTime == e.count {
			logger.Debug("StygianShadow trigger skill %d", i)
			e.currentSkill = s.Info
			e.currentSkill.Init(e.manager, e.id)
			break
		}
	}

	endTime := e.timeline[len(e.timeline)-1].TriggerTime + 1
	if e.count == endTime {
		// managerへ終了を通知
		logger.Debug("Game finished")
		e.manager.SetEnd()
	}

	return false
}

func (e *StygianShadow) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       e.id,
		Pos:      e.pos,
		IsPlayer: false,
	}
}

func (e *StygianShadow) HandleDamage(dm models.Damage) {
	logger.Debug("StygianShadow got damage %d", dm.Power)
}

func (e *StygianShadow) drawCastBar() {
	if e.currentSkill == nil {
		return
	}

	pm := e.currentSkill.GetParam()
	castTime := pm.CastTime - e.currentSkill.GetCount()
	if e.currentSkill.GetCount() != 0 && castTime > 0 {
		size := 200
		px := config.ScreenSizeX*3/4 + 50 - size/2
		py := 50
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / pm.CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "%s", pm.Name)
	}
}
