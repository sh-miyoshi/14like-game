package skill

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
)

const (
	breakCastTime = 120
)

type Break struct {
	count   int
	ownerID string
	manager models.Manager
	imgEye  int
	isLeft  bool
}

func (a *Break) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	a.imgEye = dxlib.LoadGraph("data/images/eye.png")
	if a.imgEye == -1 {
		system.FailWithError("Failed to load image")
	}
	a.isLeft = rand.Intn(2) == 0
}

func (a *Break) End() {
	dxlib.DeleteGraph(a.imgEye)
}

func (a *Break) Draw() {
	// 本体
	dxlib.DrawRotaGraph(config.ScreenSizeX/2, 180, 0.3, 0.0, a.imgEye, true)

	// 外周
	dxlib.DrawCircle(config.ScreenSizeX/2, 500, 30, 0xFFFFFF, true)
	dxlib.DrawRotaGraph(config.ScreenSizeX/2, 480, 0.3, 0.0, a.imgEye, true)

	if a.isLeft {
		dxlib.DrawCircle(200, 300, 30, 0xFFFFFF, true)
		dxlib.DrawRotaGraph(220, 300, 0.3, math.Pi/2, a.imgEye, true)
	} else {
		dxlib.DrawCircle(600, 300, 30, 0xFFFFFF, true)
		dxlib.DrawRotaGraph(580, 300, 0.3, math.Pi/2, a.imgEye, true)
	}
}

func (a *Break) Update() bool {
	a.count++
	if a.count == breakCastTime {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
		for _, obj := range objs {
			correct := math.Pi / 2
			if a.isLeft {
				correct = math.Pi * 3 / 2
			}
			if obj.Direct < correct-0.1 || obj.Direct > correct+0.1 {
				a.manager.AddDamage(models.Damage{
					ID:         uuid.New().String(),
					Power:      1,
					DamageType: models.DamageTypeObject,
					TargetID:   obj.ID,
				})
			}
		}
		return true
	}
	return false
}

func (a *Break) GetCount() int {
	return a.count
}

func (a *Break) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: breakCastTime,
		Name:     "ブレクジャ",
	}
}
