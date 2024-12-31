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
	ObjectInstPlayer int = iota
	ObjectInstCloudOfDarkness
	ObjectInstWaveGunAttacker
	ObjectInstGrimEmbraceAttacker
)

type ObjectFilter struct {
	ID   string
	Type int
}

type SkillParam struct {
	CastTime int
	Name     string
}

type ObjectParam struct {
	ID       string
	Pos      point.Point
	IsPlayer bool
	Direct   float64
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

type Skill interface {
	Init(manager Manager, ownerID string)
	End()
	Draw()
	Update() bool
	GetCount() int
	GetParam() SkillParam
}

/*
===Object===
package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type Object struct {
	id      string
	pos     point.Point
	manager models.Manager
}

func (p *Object) Init(pm interface{}, manager models.Manager) {
	p.id = uuid.New().String()
	p.manager = manager
}

func (p *Object) Draw() {
}

func (p *Object) Update() bool {
	return false
}

func (p *Object) HandleDamage(dm models.Damage) {
}

func (p *Object) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       p.id,
		Pos:      p.pos,
		IsPlayer: false,
	}
}

===Skill===
package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

const (
	attackCastTime = 120
)

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

func (a *Attack) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: attackCastTime,
		Name:     "Attack",
	}
}
*/
