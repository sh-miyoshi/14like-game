package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/enemy"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	waveGunAttackerCastTime = 240
)

type WaveGunAttacker struct {
	id        string
	pos       point.Point
	direct    int
	startTime int
	count     int
	manager   models.Manager
}

func (p *WaveGunAttacker) Init(pm interface{}, manager models.Manager) {
	p.id = uuid.New().String()
	p.manager = manager
	parsedParam := pm.(*skill.WaveGunAttackerParam)
	p.pos = parsedParam.Pos
	p.direct = parsedParam.Direct
	p.startTime = parsedParam.StartTime
}

func (p *WaveGunAttacker) Draw() {
	if p.count < p.startTime {
		dxlib.DrawCircle(p.pos.X, p.pos.Y, 10, 0xFFFFFF, true)
	} else {
		dxlib.DrawCircle(p.pos.X, p.pos.Y, 10, dxlib.GetColor(168, 88, 168), true)
	}
}

func (p *WaveGunAttacker) Update() bool {
	p.count++

	if p.count >= p.startTime+waveGunAttackerCastTime-30 {
		// 範囲
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		w := 65 // 50 * sqrt(2)
		b := p.pos
		s := point.Point{X: p.pos.X - w/2, Y: p.pos.Y}
		l := 352 // 250 * sqrt(2)
		angle := math.Pi * 3 / 4
		if p.direct == skill.WaveGunAttackRight {
			angle = -math.Pi * 3 / 4
		}

		p1 := math.Rotate(b, s, angle)
		p2 := math.Rotate(b, point.Point{X: s.X + w, Y: s.Y}, angle)
		p3 := math.Rotate(b, point.Point{X: s.X + w, Y: s.Y + l}, angle)
		p4 := math.Rotate(b, point.Point{X: s.X, Y: s.Y + l}, angle)
		dxlib.DrawQuadrangle(p1.X, p1.Y, p2.X, p2.Y, p3.X, p3.Y, p4.X, p4.Y, dxlib.GetColor(255, 255, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}
	return false
}

func (p *WaveGunAttacker) HandleDamage(dm models.Damage) {
}

func (p *WaveGunAttacker) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       p.id,
		Pos:      p.pos,
		IsPlayer: false,
	}
}
