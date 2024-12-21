package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
)

type DamageManager struct {
	damages    []models.Damage
	objManager ObjectManager
}

func (m *DamageManager) SetManager(objManager ObjectManager) {
	m.objManager = objManager
}

func (m *DamageManager) AddDamage(damage models.Damage) {
	m.damages = append(m.damages, damage)
}

func (m *DamageManager) Update() {
	for i := 0; i < len(m.damages); i++ {
		if m.damages[i].DamageType == models.TypeObject {
			// 対象のObjectにダメージを追加

			obj := m.objManager.Find(m.damages[i].TargetID)
			if obj != nil {
				obj.HandleDamage(m.damages[i].Power)
			}
		} else if m.damages[i].DamageType == models.TypeAreaCircle {
			// 範囲内のObjectにダメージを追加

		}
		// WIP: それ以外なら範囲内のObjectにダメージを追加

		m.damages = append(m.damages[:i], m.damages[i+1:]...)
		i--
	}
}
