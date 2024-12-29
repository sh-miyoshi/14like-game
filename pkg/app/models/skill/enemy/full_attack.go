package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
)

type FullAttack struct {
	Name     string
	CastTime int
	Power    int

	count   int
	ownerID string
	manager models.Manager
}

func (a *FullAttack) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *FullAttack) End() {
}

func (a *FullAttack) Draw() {
}

func (a *FullAttack) Update() bool {
	// 詠唱
	if a.count >= a.CastTime {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{Type: models.FilterObjectTypePlayer})
		for _, obj := range objs {
			a.manager.AddDamage(models.Damage{
				ID:         uuid.New().String(),
				Power:      a.Power,
				DamageType: models.DamageTypeObject,
				TargetID:   obj.ID,
			})
		}
		return true
	}

	a.count++
	return false
}

func (a *FullAttack) GetCount() int {
	return a.count
}

func (a *FullAttack) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: a.CastTime,
		Name:     a.Name,
	}
}
