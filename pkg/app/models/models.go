package models

import "github.com/sh-miyoshi/14like-game/pkg/utils/point"

const (
	FilterObjectTypeAny = iota
	FilterObjectTypePlayer
	FilterObjectTypeEnemy
)

const (
	DamageTypeObject int = iota
	DamageTypeAreaCircle
	DamageTypeAreaRect
)

type ObjectFilter struct {
	ID   string
	Type int
}

type EnemySkillParam struct {
	CastTime int
	Name     string
}

type PlayerSkillParam struct {
	CastTime   int
	RecastTime int
	Power      int
	Range      int
}

type ObjectParam struct {
	ID       string
	Pos      point.Point
	IsPlayer bool
}

type Damage struct {
	ID         string
	Power      int
	DamageType int

	// DamageTypeがTypeObjectの時使うパラメータ
	TargetID string

	// DamageTypeがTypeAreaCircleの時使うパラメータ
	CenterPos point.Point
	Range     int

	// DamageTypeがTypeAreaRectの時使うパラメータ
	RectPos     [2]point.Point
	RotateBase  point.Point
	RotateAngle float64
}
type Manager interface {
	AddDamage(damage Damage)
	GetObjectParams(filter *ObjectFilter) []ObjectParam
}

type Buff interface {
	Init(manager Manager, ownerID string)
	End()
	Update() bool
	GetIcon() int
	GetCount() int
	StackCount() int
}

type Object interface {
	Draw()
	Update()

	HandleDamage(power int)
	GetParam() ObjectParam
}

type EnemySkill interface {
	Init(manager Manager, ownerID string)
	End()
	Draw()
	Update() bool
	GetCount() int
	GetParam() EnemySkillParam
}

type PlayerSkill interface {
	Init()
	End()
	Exec(manager Manager)

	GetParam() PlayerSkillParam
	GetIcon() int
}
