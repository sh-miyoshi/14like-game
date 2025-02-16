package object

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

type StygianShadow struct {
	id      string
	manager models.Manager
	image   int
}

func (e *StygianShadow) Init(manager models.Manager) {
	e.id = uuid.New().String()
	e.manager = manager
	e.image = dxlib.LoadGraph("data/images/objects/stygian_shadow.png")
	if e.image == -1 {
		system.FailWithError("Failed to load stygian_shadow image")
	}
}

func (e *StygianShadow) End() {
}

func (e *StygianShadow) Draw() {
}

func (e *StygianShadow) Update() bool {
	return false
}
