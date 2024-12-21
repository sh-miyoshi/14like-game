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
	m.damageMgr.SetInsts(objs)
	m.objectMgr.SetObjects(objs)
}

func (m *Manager) AddDamage(damage models.Damage) {
	m.damageMgr.AddDamage(damage)
}

func (m *Manager) GetPosList() []point.Point {
	return m.objectMgr.GetPosList()
}

func (m *Manager) Update() {
	m.damageMgr.Update()
}
