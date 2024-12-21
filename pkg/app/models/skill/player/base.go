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
	Exec(manager models.Manager)

	GetParam() Param
	GetIcon() int
}
