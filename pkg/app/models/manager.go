package models

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

type Manager interface {
	AddDamage(damage Damage)
	GetPosList() []point.Point
}
