package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type Attack1 struct {
	iconImage int
}

func (a *Attack1) Init() {
	a.iconImage = dxlib.LoadGraph("data/images/attack.png")
	if a.iconImage == -1 {
		system.FailWithError("Failed to load attack image")
	}
}

func (a *Attack1) End() {
	dxlib.DeleteGraph(a.iconImage)
}

func (a *Attack1) Exec(AddDamage func(models.Damage)) {
	AddDamage(models.Damage{
		ID:         uuid.New().String(),
		Power:      a.GetParam().Power,
		DamageType: models.TypeObject,
		TargetType: models.TargetEnemy,
	})
}

func (a *Attack1) GetParam() Param {
	return Param{
		CastTime:   0,
		RecastTime: 180,
		Power:      30,
		Range:      50,
	}
}

func (a *Attack1) GetIcon() int {
	return a.iconImage
}
