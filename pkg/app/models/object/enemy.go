package object

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	Enemy1HitRange = 150
)

type Enemy1 struct {
	pos point.Point
}

func (e *Enemy1) Init() {
	e.pos.X = config.ScreenSizeX * 3 / 4
	e.pos.Y = config.ScreenSizeY / 2
}

func (e *Enemy1) End() {
}

func (e *Enemy1) Draw() {
	dxlib.DrawCircle(e.pos.X, e.pos.Y, Enemy1HitRange, dxlib.GetColor(255, 255, 255), false)
}

func (e *Enemy1) Update() {
}

func (e *Enemy1) GetPos() point.Point {
	return e.pos
}
