package manager

import (
	"fmt"

	manager "github.com/sh-miyoshi/14like-game/pkg/app/manager/internal"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
)

type Manager struct {
	isEnd     bool
	result    models.ResultInfo
	damageMgr manager.DamageManager
	objectMgr manager.ObjectManager
}

func (m *Manager) Init() {
	m.isEnd = false
	m.damageMgr.Init()
	m.objectMgr.Init()
	m.damageMgr.SetManager(&m.objectMgr)
}

func (m *Manager) AddObject(objType int, pm interface{}) string {
	var obj models.Object
	switch objType {
	case models.ObjectInstPlayer:
		tmp := &object.Player{}
		tmp.Init(m)
		obj = tmp
	case models.ObjectInstCloudOfDarkness:
		tmp := &object.CloudOfDarkness{}
		tmp.Init(m)
		obj = tmp
	case models.ObjectInstWaveGunAttacker:
		tmp := &object.WaveGunAttacker{}
		tmp.Init(pm, m)
		obj = tmp
	case models.ObjectInstGrimEmbraceAttacker:
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

func (m *Manager) SetEnd() {
	m.isEnd = true
}

func (m *Manager) IsEnd() bool {
	return m.isEnd
}

func (m *Manager) SetResult(info models.ResultInfo) {
	m.result = info
}

func (m *Manager) GetResult() models.ResultInfo {
	return m.result
}
