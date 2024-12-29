package manager

import (
	manager "github.com/sh-miyoshi/14like-game/pkg/app/manager/internal"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
)

type Manager struct {
	damageMgr manager.DamageManager
	objectMgr manager.ObjectManager
}

func (m *Manager) Init() {
	m.damageMgr.SetManager(&m.objectMgr)
}

func (m *Manager) AddObject(objType int, pm interface{}) {
	switch objType {
	case models.ObjectTypePlayer:
		player := &object.Player{}
		player.Init(m)
		m.objectMgr.AddObject(player)
	case models.ObjectTypeEnemy:
		enemy1 := &object.Enemy1{}
		enemy1.Init(m)
		m.objectMgr.AddObject(enemy1)
	case models.ObjectTypeBombBoulder:
		obj := &object.BombBoulder{}
		obj.Init(pm)
		m.objectMgr.AddObject(obj)
	}
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
