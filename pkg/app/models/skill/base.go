package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Param struct {
	CastTime   int
	RecastTime int
	Power      int
	Range      int
}

type Skill interface {
	Init()
	End()
	Exec(AddDamage func(models.Damage))

	GetParam() Param
	GetIcon() int
}
