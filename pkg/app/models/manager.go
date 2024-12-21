package models

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

type ObjectFilter struct {
	ID string
}

type Manager interface {
	AddDamage(damage Damage)
	GetPosList(filter *ObjectFilter) []point.Point
}
