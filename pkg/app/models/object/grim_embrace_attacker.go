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
	id      string
	pos     point.Point
	manager models.Manager
	direct  float64
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
	b := p.pos
	s := b
	width := 30
	length := 100
	p1 := math.Rotate(b, s, p.direct)
	p2 := math.Rotate(b, point.Point{X: s.X + width, Y: s.Y}, p.direct)
	p3 := math.Rotate(b, point.Point{X: s.X + width, Y: s.Y + length}, p.direct)
	p4 := math.Rotate(b, point.Point{X: s.X, Y: s.Y + length}, p.direct)
	dxlib.DrawQuadrangle(
		p1.X, p1.Y,
		p2.X, p2.Y,
		p3.X, p3.Y,
		p4.X, p4.Y,
		dxlib.GetColor(255, 255, 0),
		true,
	)
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (p *GrimEmbraceAttacker) Update() bool {
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
