package buff

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/system"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
)

const (
	PoisonCount = 15 * 60 // 15秒
	PoisonPower = 10
)

type Poison struct {
	icon    int
	count   int
	manager models.Manager
	ownerID string
}

func (p *Poison) Init(manager models.Manager, ownerID string) {
	p.icon = dxlib.LoadGraph("data/images/buff/poison.png")
	p.count = PoisonCount
	if p.icon == -1 {
		system.FailWithError("Failed to load poison buff image")
	}
	p.manager = manager
	p.ownerID = ownerID
}

func (p *Poison) End() {
	dxlib.DeleteGraph(p.icon)
}

func (p *Poison) Update() bool {
	if p.count%60 == 0 {
		p.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      PoisonPower,
			DamageType: models.TypeObject,
			TargetID:   p.ownerID,
		})
	}

	p.count--
	return p.count <= 0
}

func (p *Poison) GetIcon() int {
	return p.icon
}

func (p *Poison) GetCount() int {
	return p.count
}

func (p *Poison) StackCount() int {
	return 0
}
