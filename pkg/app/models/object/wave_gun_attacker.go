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
	waveGunAttackerDamageTime = 300
	waveGunAttackerEndTime    = waveGunAttackerDamageTime + 50
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

	width := 65 // 50 * sqrt(2)
	base := p.pos
	start := point.Point{X: p.pos.X - width/2, Y: p.pos.Y}
	length := 352 // 250 * sqrt(2)
	angle := math.Pi * 3 / 4
	if p.direct == skill.WaveGunAttackRight {
		angle = -math.Pi * 3 / 4
	}

	if p.count == p.startTime+waveGunAttackerDamageTime {
		p.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      1,
			DamageType: models.DamageTypeAreaRect,
			RectPos: [2]point.Point{
				start,
				{X: start.X + width, Y: start.Y + length},
			},
			RotateBase:  base,
			RotateAngle: angle,
		})
	}
	if p.count >= p.startTime+waveGunAttackerDamageTime {
		// 範囲
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		p1 := math.Rotate(base, start, angle)
		p2 := math.Rotate(base, point.Point{X: start.X + width, Y: start.Y}, angle)
		p3 := math.Rotate(base, point.Point{X: start.X + width, Y: start.Y + length}, angle)
		p4 := math.Rotate(base, point.Point{X: start.X, Y: start.Y + length}, angle)
		dxlib.DrawQuadrangle(p1.X, p1.Y, p2.X, p2.Y, p3.X, p3.Y, p4.X, p4.Y, dxlib.GetColor(255, 255, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}

	if p.count == p.startTime+waveGunAttackerEndTime {
		return true
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
