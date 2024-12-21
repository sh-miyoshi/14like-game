package manager

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models/object"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type ObjectManager struct {
	objects []object.Object
}

func (m *ObjectManager) SetObjects(objs []object.Object) {
	m.objects = objs
}

func (m *ObjectManager) GetPosList() []point.Point {
	res := []point.Point{}
	for _, o := range m.objects {
		res = append(res, o.GetPos())
	}
	return res
}
