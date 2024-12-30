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

const (
	ObjectTypePlayer int = iota
	ObjectTypeEnemy
	ObjectTypeBombBoulder
	ObjectTypeNonAttackPlayer
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
	Buffs      []Buff

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
	AddObject(objType int, pm interface{}) string
	AddDamage(damage Damage)
	GetObjectParams(filter *ObjectFilter) []ObjectParam
	GetObjects(filter *ObjectFilter) []Object
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
	Update() bool

	HandleDamage(dm Damage)
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

/*
===EnemySkill===
package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Attack struct {
	count   int
	ownerID string
	manager models.Manager
}

func (a *Attack) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *Attack) End() {
}

func (a *Attack) Draw() {
}

func (a *Attack) Update() bool {
	a.count++
	return false
}

func (a *Attack) GetCount() int {
	return a.count
}

func (a *Attack) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: 10,
		Name:     "Attack",
	}
}
===
*/
