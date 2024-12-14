package object

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	PlayerHitRange = 50
)

type Player struct {
	pos point.Point
}

func (p *Player) Init() {
	p.pos.X = config.ScreenSizeY / 4
	p.pos.Y = config.ScreenSizeY / 2
}

func (p *Player) End() {
}

func (p *Player) Draw() {
	dxlib.DrawCircle(p.pos.X, p.pos.Y, PlayerHitRange, dxlib.GetColor(255, 255, 255), false)
}

func (p *Player) Update() {
}
