package object

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/skill"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

var (
	debug = false
)

type CloudOfDarkness struct {
	id           string
	pos          point.Point
	timeline     []SkillTimeline
	count        int
	currentSkill models.Skill
	manager      models.Manager
	image        int
}

func (e *CloudOfDarkness) Init(manager models.Manager) {
	e.id = uuid.New().String()
	e.pos.X = config.ScreenSizeX / 2
	e.pos.Y = 50
	e.manager = manager
	e.currentSkill = nil
	e.image = dxlib.LoadGraph("data/images/objects/cloud_of_darkness.png")
	if e.image == -1 {
		system.FailWithError("Failed to load cloud_of_darkness image")
	}

	if debug {
		// デバッグ
		e.timeline = []SkillTimeline{
			{60, &skill.OnlyCast{CastTime: 120, Name: "闇の大氾濫", Text: "おわり～"}},
		}
	} else {
		// パターン1
		e.timeline = []SkillTimeline{
			{180, &skill.GrimEmbrace{}},
			{300, &skill.WaveGun{}},
		}
		n := rand.Intn(4)
		if n%2 == 0 {
			e.timeline = append(e.timeline, SkillTimeline{660, &skill.Aero{}})
		} else {
			e.timeline = append(e.timeline, SkillTimeline{660, &skill.Death{}})
		}
		if n/2 == 0 {
			e.timeline = append(e.timeline, SkillTimeline{800, &skill.OnlyCast{CastTime: 240, Name: "エンエアロジャ"}})
		} else {
			e.timeline = append(e.timeline, SkillTimeline{800, &skill.OnlyCast{CastTime: 240, Name: "エンデスジャ"}})
		}
		e.timeline = append(e.timeline, []SkillTimeline{
			{1020, &skill.RapidParticleBeam{}},
			{1080, &skill.WaveGun{}},
			{1440, &skill.BladeOfDarkness{AttackType: skill.BladeOfDarknessAttackLeft}},
		}...)
		if n/2 == 0 {
			e.timeline = append(e.timeline, SkillTimeline{1470, &skill.Aero{CastTime: 60}})
		} else {
			e.timeline = append(e.timeline, SkillTimeline{1470, &skill.Death{CastTime: 60}})
		}
		e.timeline = append(e.timeline, []SkillTimeline{
			{1830, &skill.OnlyCast{CastTime: 240, Name: "フレア", Text: "本番は散会～"}},
			{1890, &skill.OnlyCast{CastTime: 240, Name: "闇の大氾濫", Text: "おわり～"}},
		}...)

		// WIP: パターン2
	}
}

func (e *CloudOfDarkness) End() {
}

func (e *CloudOfDarkness) Draw() {
	dxlib.DrawRotaGraph(e.pos.X, e.pos.Y, 0.5, 0.0, e.image, true)
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
		if s.TriggerTime == e.count {
			logger.Debug("CloudOfDarkness trigger skill %d", i)
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
		px := config.ScreenSizeX*3/4 + 50 - size/2
		py := 50
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / pm.CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "%s", pm.Name)
	}
}
