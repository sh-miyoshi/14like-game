package models

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

const (
	FilterObjectTypeAny = iota
	FilterObjectTypePlayer
	FilterObjectTypeEnemy
)

type ObjectFilter struct {
	ID   string
	Type int
}

type Manager interface {
	AddDamage(damage Damage)
	GetPosList(filter *ObjectFilter) []point.Point
}
