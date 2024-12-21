package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
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
			// WIP: 現状は敵の攻撃のみサポート
			// 範囲内のObjectにダメージを追加
			objs := m.objManager.GetObjects(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
			for _, obj := range objs {
				x1 := obj.GetParam().Pos.X
				y1 := obj.GetParam().Pos.Y
				x2 := m.damages[i].CenterPos.X
				y2 := m.damages[i].CenterPos.Y

				dist2 := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
				hitRange := config.PlayerHitRange + m.damages[i].Range

				if dist2 < hitRange*hitRange {
					obj.HandleDamage(m.damages[i].Power)
				}
			}
		}

		m.damages = append(m.damages[:i], m.damages[i+1:]...)
		i--
	}
}
