package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
)

type DamageManager struct {
	damages  []models.Damage
	objInsts []object.Object
}

func (m *DamageManager) SetInsts(objs []object.Object) {
	m.objInsts = objs
}

func (m *DamageManager) AddDamage(damage models.Damage) {
	m.damages = append(m.damages, damage)
}

func (m *DamageManager) Update() {
	for i := 0; i < len(m.damages); i++ {
		if m.damages[i].DamageType == models.TypeObject {
			// 対象のObjectにダメージを追加
			obj := m.findObject(m.damages[i].TargetType)
			if obj != nil {
				obj.HandleDamage(m.damages[i].Power)
			}
		}
		// WIP: それ以外なら範囲内のObjectにダメージを追加

		m.damages = append(m.damages[:i], m.damages[i+1:]...)
		i--
	}
}

func (m *DamageManager) findObject(typ int) object.Object {
	for _, obj := range m.objInsts {
		if typ == models.TargetPlayer && obj.IsPlayer() {
			return obj
		} else if typ == models.TargetEnemy && !obj.IsPlayer() {
			return obj
		}
	}
	return nil
}
