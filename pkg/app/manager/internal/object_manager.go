package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type ObjectManager struct {
	objects []models.Object
}

func (m *ObjectManager) SetObjects(objs []models.Object) {
	m.objects = objs
}

func (m *ObjectManager) GetObjects(filter *models.ObjectFilter) []models.Object {
	res := []models.Object{}
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
		res = append(res, o)
	}
	return res
}

func (m *ObjectManager) GetPosList(filter *models.ObjectFilter) []point.Point {
	res := []point.Point{}
	objs := m.GetObjects(filter)
	for _, o := range objs {
		res = append(res, o.GetParam().Pos)
	}
	return res
}

func (m *ObjectManager) GetObjectsID(filter *models.ObjectFilter) []string {
	res := []string{}
	objs := m.GetObjects(filter)
	for _, o := range objs {
		res = append(res, o.GetParam().ID)
	}
	return res
}

func (m *ObjectManager) Find(id string) models.Object {
	for _, o := range m.objects {
		if o.GetParam().ID == id {
			return o
		}
	}
	return nil
}
