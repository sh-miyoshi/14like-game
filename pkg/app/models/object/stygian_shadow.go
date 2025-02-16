package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type StygianShadow struct {
	id      string
	manager models.Manager
	image   int
	pos     point.Point
}

func (e *StygianShadow) Init(manager models.Manager) {
	e.id = uuid.New().String()
	e.manager = manager
	e.image = dxlib.LoadGraph("data/images/objects/stygian_shadow.png")
	if e.image == -1 {
		system.FailWithError("Failed to load stygian_shadow image")
	}
	e.pos = point.Point{X: config.ScreenSizeX / 2, Y: 200}
}

func (e *StygianShadow) End() {
}

func (e *StygianShadow) Draw() {
	dxlib.DrawRotaGraph(e.pos.X, e.pos.Y, 0.3, 0.0, e.image, true)
}

func (e *StygianShadow) Update() bool {
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
