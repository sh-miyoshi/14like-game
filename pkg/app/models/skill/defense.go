package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type Defense1 struct {
	iconImage int
}

func (d *Defense1) Init() {
	d.iconImage = dxlib.LoadGraph("data/images/defense.png")
	if d.iconImage == -1 {
		system.FailWithError("Failed to load defense image")
	}
}

func (d *Defense1) End() {
	dxlib.DeleteGraph(d.iconImage)
}

func (d *Defense1) Exec(AddDamage func(models.Damage)) {
	// WIP: 自分の防御力をアップ or 味方全体の防御力をアップ
}

func (d *Defense1) GetParam() Param {
	return Param{
		CastTime:   0,
		RecastTime: 540,
		Power:      30,
		Range:      -1,
	}
}

func (d *Defense1) GetIcon() int {
	return d.iconImage
}
