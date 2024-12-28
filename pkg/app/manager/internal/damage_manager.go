package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
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
		switch m.damages[i].DamageType {
		case models.TypeObject:
			// 対象のObjectにダメージを追加
			obj := m.objManager.Find(m.damages[i].TargetID)
			if obj != nil {
				obj.HandleDamage(m.damages[i].Power)
			}
		case models.TypeAreaCircle, models.TypeAreaRect:
			// WIP: 現状は敵の攻撃のみサポート
			// 範囲内のObjectにダメージを追加
			objs := m.objManager.GetObjects(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
			for _, obj := range objs {
				if m.damages[i].DamageType == models.TypeAreaCircle {
					hitRange := config.PlayerHitRange + m.damages[i].Range
					if isCircleHit(obj.GetParam().Pos, m.damages[i].CenterPos, hitRange) {
						obj.HandleDamage(m.damages[i].Power)
					}
				} else if m.damages[i].DamageType == models.TypeAreaRect {
					pos := math.Rotate(m.damages[i].RotateBase, obj.GetParam().Pos, -m.damages[i].RotateAngle)
					if isRectHit(pos, config.PlayerHitRange, m.damages[i].RectPos) {
						obj.HandleDamage(m.damages[i].Power)
					}
				}
			}
		}

		m.damages = append(m.damages[:i], m.damages[i+1:]...)
		i--
	}
}

func isCircleHit(p1, p2 point.Point, r int) bool {
	x1 := p1.X
	y1 := p1.Y
	x2 := p2.X
	y2 := p2.Y

	dist2 := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)
	return dist2 < r*r
}

func isRectHit(p point.Point, r int, rectPos [2]point.Point) bool {
	xp := p.X
	yp := p.Y
	x1 := rectPos[0].X
	y1 := rectPos[0].Y
	x2 := rectPos[1].X
	y2 := rectPos[1].Y

	if xp > x1 && xp < x2 && yp > (y1-r) && yp < (y2+r) {
		return true
	}

	if xp > (x1-r) && xp < (x2+r) && yp > y1 && yp < y2 {
		return true
	}

	r2 := r * r
	if (x1-xp)*(x1-xp)+(y1-yp)*(y1-yp) < r2 {
		return true
	}
	if (x2-xp)*(x2-xp)+(y1-yp)*(y1-yp) < r2 {
		return true
	}
	if (x2-xp)*(x2-xp)+(y2-yp)*(y2-yp) < r2 {
		return true
	}
	if (x1-xp)*(x1-xp)+(y2-yp)*(y2-yp) < r2 {
		return true
	}

	return false
}
