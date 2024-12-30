package skill

import (
	"fmt"
	"math/rand"

	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	bombBoulderInitDelay = 45
)

type BombBoulderParam struct {
	StartTime int
	Pos       point.Point
}

type BombBoulderMgr struct {
	ownerID   string
	manager   models.Manager
	attackIDs string
}

func (a *BombBoulderMgr) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	yindex := [3]int{1, 0, 2}
	if rand.Intn(2) == 0 {
		yindex = [3]int{1, 2, 0}
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			id := a.manager.AddObject(models.ObjectTypeBombBoulder, &BombBoulderParam{
				Pos: point.Point{
					X: (config.ScreenSizeX/4+40)*(x+1) - 40,
					Y: config.ScreenSizeY/4*(yindex[y]+1) - 20,
				},
				StartTime: bombBoulderInitDelay * y,
			})
			a.attackIDs = fmt.Sprintf("%s,%s", a.attackIDs, id)
		}
	}
	logger.Debug("bomb boulders: %+v", a.attackIDs)
}

func (a *BombBoulderMgr) End() {
}

func (a *BombBoulderMgr) Draw() {
}

func (a *BombBoulderMgr) Update() bool {
	// すべての処理が終わっていたら終了
	objs := a.manager.GetObjects(&models.ObjectFilter{ID: a.attackIDs})
	return len(objs) == 0
}

func (a *BombBoulderMgr) GetCount() int {
	return 0
}

func (a *BombBoulderMgr) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: 0,
		Name:     "ボムボルダー",
	}
}
