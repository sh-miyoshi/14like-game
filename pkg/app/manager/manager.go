package manager

import (
	manager "github.com/sh-miyoshi/14like-game/pkg/app/manager/internal"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type Manager struct {
	damageMgr manager.DamageManager
	objectMgr manager.ObjectManager
}

func (m *Manager) SetObjects(objs []object.Object) {
	m.objectMgr.SetObjects(objs)
	m.damageMgr.SetManager(m.objectMgr)
}

func (m *Manager) AddDamage(damage models.Damage) {
	m.damageMgr.AddDamage(damage)
}

func (m *Manager) GetPosList(filter *models.ObjectFilter) []point.Point {
	return m.objectMgr.GetPosList(filter)
}

func (m *Manager) GetObjectsID(filter *models.ObjectFilter) []string {
	return m.objectMgr.GetObjectsID(filter)
}

func (m *Manager) Update() {
	m.damageMgr.Update()
}
