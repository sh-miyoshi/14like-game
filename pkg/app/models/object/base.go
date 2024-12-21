package object

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

type Param struct {
	ID       string
	Pos      point.Point
	IsPlayer bool
}

type Object interface {
	Draw()
	Update()

	HandleDamage(power int)
	GetParam() Param
}
