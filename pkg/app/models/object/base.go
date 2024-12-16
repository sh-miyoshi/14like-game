package object

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

type Object interface {
	Draw()
	Update()

	GetPos() point.Point
	IsPlayer() bool
	HandleDamage(power int)
}
