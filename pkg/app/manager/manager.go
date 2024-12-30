package manager

import (
	"fmt"

	manager "github.com/sh-miyoshi/14like-game/pkg/app/manager/internal"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
)

type Manager struct {
	damageMgr manager.DamageManager
	objectMgr manager.ObjectManager
}

func (m *Manager) Init() {
	m.damageMgr.SetManager(&m.objectMgr)
}

func (m *Manager) AddObject(objType int, pm interface{}) string {
	var obj models.Object
	switch objType {
	case models.ObjectTypePlayer:
		tmp := &object.Player{}
		tmp.Init(m)
		obj = tmp
	case models.ObjectTypeEnemy:
		tmp := &object.Enemy1{}
		tmp.Init(m)
		obj = tmp
	case models.ObjectTypeBombBoulder:
		tmp := &object.BombBoulder{}
		tmp.Init(pm, m)
		obj = tmp
	case models.ObjectTypeNonAttackPlayer:
		tmp := &object.NonAttackPlayer{}
		tmp.Init(m)
		obj = tmp
	case models.ObjectTypeCloudOfDarkness:
		tmp := &object.CloudOfDarkness{}
		tmp.Init(m)
		obj = tmp
	case models.ObjectTypeWaveGunAttacker:
		tmp := &object.WaveGunAttacker{}
		tmp.Init(pm, m)
		obj = tmp
	case models.ObjectTypeGrimEmbraceAttacker:
		tmp := &object.GrimEmbraceAttacker{}
		tmp.Init(pm, m)
		obj = tmp
	default:
		system.FailWithError(fmt.Sprintf("Unknown object type %d", objType))
	}
	m.objectMgr.AddObject(obj)
	return obj.GetParam().ID
}

func (m *Manager) AddDamage(damage models.Damage) {
	m.damageMgr.AddDamage(damage)
}

func (m *Manager) GetObjectParams(filter *models.ObjectFilter) []models.ObjectParam {
	return m.objectMgr.GetObjectParams(filter)
}

func (m *Manager) GetObjects(filter *models.ObjectFilter) []models.Object {
	return m.objectMgr.GetObjects(filter)
}

func (m *Manager) Draw() {
	m.objectMgr.Draw()
}

func (m *Manager) Update() {
	m.damageMgr.Update()
	m.objectMgr.Update()
}
