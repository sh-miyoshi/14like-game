package object

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

type Object interface {
	Init()
	End()

	Draw()
	Update()

	GetPos() point.Point
}
