package models

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

const (
	TargetPlayer int = iota
	TargetEnemy
)

const (
	TypeObject int = iota
	TypeArea
)

type Damage struct {
	ID         string
	Power      int
	DamageType int

	// DamageTypeがTypeObjectの時使うパラメータ
	TargetType int

	// DamageTypeがTypeAreaの時使うパラメータ
	CenterPos point.Point
	Range     int
	TTL       int
}
