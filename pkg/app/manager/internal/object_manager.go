package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type ObjectManager struct {
	objects []object.Object
}

func (m *ObjectManager) SetObjects(objs []object.Object) {
	m.objects = objs
}

func (m *ObjectManager) GetPosList(filter *models.ObjectFilter) []point.Point {
	res := []point.Point{}
	for _, o := range m.objects {
		if filter != nil {
			if filter.ID != "" && filter.ID != o.GetParam().ID {
				continue
			}
			switch filter.Type {
			case models.FilterObjectTypePlayer:
				if !o.GetParam().IsPlayer {
					continue
				}
			case models.FilterObjectTypeEnemy:
				if o.GetParam().IsPlayer {
					continue
				}
			}
		}
		res = append(res, o.GetParam().Pos)
	}
	return res
}
