package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/buff"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	grimEmbraceCastTime = 180
)

type GrimEmbrace struct {
	count          int
	ownerID        string
	manager        models.Manager
	pos            point.Point
	targetObjParam models.ObjectParam
}

func (a *GrimEmbrace) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.pos = point.Point{X: config.ScreenSizeX / 2, Y: 50}
}

func (a *GrimEmbrace) End() {
}

func (a *GrimEmbrace) Draw() {
	if a.count == 0 {
		return
	}

	t := int32(3)
	dxlib.DrawLine(
		a.pos.X, a.pos.Y,
		a.targetObjParam.Pos.X, a.targetObjParam.Pos.Y,
		dxlib.GetColor(255, 0, 0),
		dxlib.DrawLineOption{
			Thickness: &t,
		},
	)
}

func (a *GrimEmbrace) Update() bool {
	objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.ObjectTypePlayer})
	if len(objs) == 0 {
		return true
	}
	a.targetObjParam = objs[0]

	a.count++
	if a.count == grimEmbraceCastTime {
		bf := &buff.GrimEmbrace{Count: 54 * 60}
		bf.Init(a.manager, a.ownerID)

		a.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      0,
			DamageType: models.DamageTypeObject,
			TargetID:   a.targetObjParam.ID,
			Buffs:      []models.Buff{bf},
		})
		return true
	}
	return false
}

func (a *GrimEmbrace) GetCount() int {
	return a.count
}

func (a *GrimEmbrace) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: grimEmbraceCastTime,
		Name:     "グリムエンブレイス",
	}
}
