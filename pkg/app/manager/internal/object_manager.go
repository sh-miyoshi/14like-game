package manager

import (
	"strings"

	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/stretchr/stew/slice"
)

type ObjectManager struct {
	objects []models.Object
}

func (m *ObjectManager) Init() {
	m.objects = []models.Object{}
}

func (m *ObjectManager) AddObject(obj models.Object) {
	m.objects = append(m.objects, obj)
}

func (m *ObjectManager) GetObjects(filter *models.ObjectFilter) []models.Object {
	res := []models.Object{}
	for _, o := range m.objects {
		if filter != nil {
			if filter.ID != "" {
				ids := strings.Split(filter.ID, ",")
				if !slice.Contains(ids, o.GetParam().ID) {
					continue
				}
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

func (m *ObjectManager) GetObjectParams(filter *models.ObjectFilter) []models.ObjectParam {
	res := []models.ObjectParam{}
	objs := m.GetObjects(filter)
	for _, o := range objs {
		res = append(res, o.GetParam())
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

func (m *ObjectManager) Draw() {
	for _, o := range m.objects {
		o.Draw()
	}
}

func (m *ObjectManager) Update() {
	for i := 0; i < len(m.objects); i++ {
		if m.objects[i].Update() {
			m.objects = append(m.objects[:i], m.objects[i+1:]...)
			i--
		}
	}
}
