package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type GrimEmbraceAttacker struct {
	id      string
	pos     point.Point
	manager models.Manager
}

func (p *GrimEmbraceAttacker) Init(pm interface{}, manager models.Manager) {
	p.id = uuid.New().String()
	p.manager = manager
}

func (p *GrimEmbraceAttacker) Draw() {
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
