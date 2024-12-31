package buff

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
)

type GrimEmbraceAttackerParam struct {
	TargetID string
	IsFront  bool
}

type GrimEmbrace struct {
	Count   int
	IsFront bool

	icon    int
	manager models.Manager
	ownerID string
}

func (p *GrimEmbrace) Init(manager models.Manager, ownerID string) {
	p.icon = dxlib.LoadGraph("data/images/buff/grim_embrace.png")
	if p.icon == -1 {
		system.FailWithError("Failed to load grim embrace buff image")
	}
	p.manager = manager
	p.ownerID = ownerID
}

func (p *GrimEmbrace) End() {
	dxlib.DeleteGraph(p.icon)
}

func (p *GrimEmbrace) Update() bool {
	p.Count--
	if p.Count == 0 {
		logger.Debug("add GrimEmbrace action")
		p.manager.AddObject(
			models.ObjectInstGrimEmbraceAttacker,
			&GrimEmbraceAttackerParam{
				TargetID: p.ownerID,
				IsFront:  p.IsFront,
			},
		)
		return true
	}
	return false
}

func (p *GrimEmbrace) GetIcon() int {
	return p.icon
}

func (p *GrimEmbrace) GetCount() int {
	return p.Count
}

func (p *GrimEmbrace) StackCount() int {
	return 0
}
