package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	skill "github.com/sh-miyoshi/14like-game/pkg/app/models/skill/enemy"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	bombBoulderCastTime = 180
	bombBoulderRange    = 180
)

type BombBoulder struct {
	id        string
	pos       point.Point
	count     int
	startTime int
	manager   models.Manager
}

func (b *BombBoulder) Init(pm interface{}, manager models.Manager) {
	b.id = uuid.New().String()
	parsedParam := pm.(*skill.BombBoulderParam)
	b.pos = parsedParam.Pos
	b.startTime = parsedParam.StartTime
	b.manager = manager
}

func (b *BombBoulder) Draw() {
	if b.count >= b.startTime {
		w := 30
		h := 50
		dxlib.DrawBox(b.pos.X-w/2, b.pos.Y, b.pos.X+w/2, b.pos.Y+h, dxlib.GetColor(255, 255, 255), true)
	}

	if b.count >= b.startTime+bombBoulderCastTime-20 {
		// 範囲
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		dxlib.DrawCircle(b.pos.X, b.pos.Y, bombBoulderRange, dxlib.GetColor(255, 255, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}
}

func (b *BombBoulder) Update() bool {
	b.count++
	if b.count == b.startTime+bombBoulderCastTime {
		b.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      120,
			DamageType: models.DamageTypeAreaCircle,
			CenterPos:  b.pos,
			Range:      bombBoulderRange,
		})
		return true
	}
	return false
}

func (b *BombBoulder) HandleDamage(dm models.Damage) {
	// WIP
}

func (b *BombBoulder) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       b.id,
		Pos:      b.pos,
		IsPlayer: false,
	}
}
