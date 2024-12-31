package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/buff"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type GrimEmbraceAttacker struct {
	id        string
	pos       point.Point
	manager   models.Manager
	direct    float64
	count     int
	attackPos [4]point.Point
}

func (p *GrimEmbraceAttacker) Init(pm interface{}, manager models.Manager) {
	p.id = uuid.New().String()
	p.manager = manager
	parsedParam := pm.(*buff.GrimEmbraceAttackerParam)
	objs := manager.GetObjectParams(&models.ObjectFilter{ID: parsedParam.TargetID})
	if len(objs) > 0 {
		p.pos = objs[0].Pos
		p.direct = objs[0].Direct
		logger.Debug("GrimEmbraceAttacker target %+v", objs[0])
	}
	if !parsedParam.IsFront {
		p.direct += math.Pi
	}
}

func (p *GrimEmbraceAttacker) Draw() {
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)

	dxlib.DrawQuadrangle(
		p.attackPos[0].X, p.attackPos[0].Y,
		p.attackPos[1].X, p.attackPos[1].Y,
		p.attackPos[2].X, p.attackPos[2].Y,
		p.attackPos[3].X, p.attackPos[3].Y,
		dxlib.GetColor(255, 255, 0),
		true,
	)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (p *GrimEmbraceAttacker) Update() bool {
	if p.count == 0 {
		width := 30
		length := 100
		b := p.pos
		s := point.Point{X: b.X - width/2, Y: b.Y}
		p.attackPos[0] = math.Rotate(b, s, p.direct)
		p.attackPos[1] = math.Rotate(b, point.Point{X: s.X + width, Y: s.Y}, p.direct)
		p.attackPos[2] = math.Rotate(b, point.Point{X: s.X + width, Y: s.Y + length}, p.direct)
		p.attackPos[3] = math.Rotate(b, point.Point{X: s.X, Y: s.Y + length}, p.direct)
	}

	if p.count > 30 {
		p.manager.AddDamage(models.Damage{
			ID:          uuid.New().String(),
			Power:       1,
			DamageType:  models.DamageTypeAreaRect,
			RectPos:     [2]point.Point{p.attackPos[0], p.attackPos[2]},
			RotateBase:  p.pos,
			RotateAngle: p.direct,
		})
		return true
	}

	p.count++
	return false
}

func (p *GrimEmbraceAttacker) HandleDamage(dm models.Damage) {
}

func (p *GrimEmbraceAttacker) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       p.id,
		Pos:      p.pos,
		IsPlayer: false,
	}
}
