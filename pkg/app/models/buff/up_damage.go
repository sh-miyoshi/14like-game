package buff

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

const (
	UpDamageCount = 60 * 60 // 60sec
)

type UpDamage struct {
	icon    int
	count   int
	manager models.Manager
	ownerID string
	stack   int
}

func (b *UpDamage) Init(manager models.Manager, ownerID string) {
	b.icon = dxlib.LoadGraph("data/images/buff/up_damage.png")
	if b.icon == -1 {
		system.FailWithError("Failed to load up_damage buff image")
	}
	b.count = UpDamageCount
	b.stack = 1
	b.manager = manager
	b.ownerID = ownerID
}

func (b *UpDamage) End() {
	dxlib.DeleteGraph(b.icon)
}

func (b *UpDamage) Update() bool {
	b.count--
	return b.count <= 0
}

func (b *UpDamage) GetIcon() int {
	return b.icon
}

func (b *UpDamage) GetCount() int {
	return b.count
}

func (b *UpDamage) StackCount() int {
	return b.stack
}

func (b *UpDamage) UpStack() {
	b.stack++
}
