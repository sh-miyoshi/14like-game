package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/enemy"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
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
	dxlib.DrawCircle(p.pos.X, p.pos.Y, 10, 0xFFFFFF, true)
	// WIP
}

func (p *WaveGunAttacker) Update() bool {
	p.count++
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
