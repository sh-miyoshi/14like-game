package manager

import (
	manager "github.com/sh-miyoshi/14like-game/pkg/app/manager/internal"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
)

type Manager struct {
	damageMgr manager.DamageManager
	objectMgr manager.ObjectManager
}

func (m *Manager) SetObjects(objs []models.Object) {
	m.objectMgr.SetObjects(objs)
	m.damageMgr.SetManager(m.objectMgr)
}

func (m *Manager) AddDamage(damage models.Damage) {
	m.damageMgr.AddDamage(damage)
}

func (m *Manager) GetObjectParams(filter *models.ObjectFilter) []models.ObjectParam {
	return m.objectMgr.GetObjectParams(filter)
}

func (m *Manager) Update() {
	m.damageMgr.Update()
}
