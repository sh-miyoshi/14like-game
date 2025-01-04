package skill

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/buff"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	grimEmbraceCastTime = 300
)

type GrimEmbrace struct {
	count          int
	ownerID        string
	manager        models.Manager
	pos            point.Point
	targetObjParam models.ObjectParam
	imgFront       int
	imgBack        int
	isFront        bool
}

func (a *GrimEmbrace) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.pos = point.Point{X: config.ScreenSizeX / 2, Y: 90}
	a.imgFront = dxlib.LoadGraph("data/images/grim_embrace_front.png")
	if a.imgFront == -1 {
		system.FailWithError("Failed to load grim_embrace_front image")
	}
	a.imgBack = dxlib.LoadGraph("data/images/grim_embrace_back.png")
	if a.imgBack == -1 {
		system.FailWithError("Failed to load grim_embrace_back image")
	}
}

func (a *GrimEmbrace) End() {
	dxlib.DeleteGraph(a.imgFront)
	dxlib.DeleteGraph(a.imgBack)
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

	if a.isFront {
		dxlib.DrawRotaGraph(a.pos.X, a.pos.Y, 1.0, 0.0, a.imgFront, true)
	} else {
		dxlib.DrawRotaGraph(a.pos.X, a.pos.Y, 1.0, 0.0, a.imgBack, true)
	}
}

func (a *GrimEmbrace) Update() bool {
	objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
	if len(objs) == 0 {
		return true
	}
	a.targetObjParam = objs[0]

	if a.count == 0 {
		a.isFront = rand.Intn(2) == 0
	}

	a.count++
	if a.count >= grimEmbraceCastTime-20 {
		if a.isFront {
			a.pos.Y -= 4
		} else {
			a.pos.Y += 4
		}
	}

	if a.count == grimEmbraceCastTime {
		tm := 22
		if rand.Intn(2) == 0 {
			tm = 54
		}

		a.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      0,
			DamageType: models.DamageTypeObject,
			TargetID:   a.targetObjParam.ID,
			Buffs:      []models.Buff{&buff.GrimEmbrace{Count: tm * 60, IsFront: a.isFront}},
		})
		return true
	}
	return false
}

func (a *GrimEmbrace) GetCount() int {
	return a.count
}

func (a *GrimEmbrace) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: grimEmbraceCastTime,
		Name:     "グリムエンブレイス",
	}
}
