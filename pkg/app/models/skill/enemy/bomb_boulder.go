package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type BombBoulder struct {
	count     int
	ownerID   string
	attackPos point.Point
	manager   models.Manager
}

func (a *BombBoulder) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *BombBoulder) End() {
}

func (a *BombBoulder) Draw() {
}

func (a *BombBoulder) Update() bool {
	a.count++
	return false
}

func (a *BombBoulder) GetCount() int {
	return a.count
}

func (a *BombBoulder) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: 0,
		Name:     "ボムボルダー",
	}
}
