package models

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

const (
	TypeObject int = iota
	TypeAreaCircle
)

type Damage struct {
	ID         string
	Power      int
	DamageType int

	// DamageTypeがTypeObjectの時使うパラメータ
	TargetID string

	// DamageTypeがTypeAreaCircleの時使うパラメータ
	CenterPos point.Point
	Range     int
}
