package skill

import (
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	bombBoulderInitDelay = 120
	bombBoulderCastTime  = 180
	bombBoulderRange     = 180
)

type bombBoulder struct {
	pos       point.Point
	count     int
	startTime int
}

type BombBoulderMgr struct {
	ownerID string
	manager models.Manager
	attacks [9]bombBoulder
}

func (a *BombBoulderMgr) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			a.attacks[y*3+x] = bombBoulder{
				pos: point.Point{
					X: (config.ScreenSizeX/4+40)*(x+1) - 40,
					Y: config.ScreenSizeY/4*(y+1) - 20,
				},
				startTime: bombBoulderInitDelay * y,
			}
		}
	}
	logger.Debug("bomb boulders: %+v", a.attacks)
}

func (a *BombBoulderMgr) End() {
}

func (a *BombBoulderMgr) Draw() {
	for _, b := range a.attacks {
		b.Draw()
	}
}

func (a *BombBoulderMgr) Update() bool {
	end := true
	for i := range a.attacks {
		if !a.attacks[i].Update() {
			end = false
		}
	}
	return end
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

func (b *bombBoulder) Draw() {
	if b.count >= b.startTime {
		w := 30
		h := 50
		dxlib.DrawBox(b.pos.X-w/2, b.pos.Y, b.pos.X+w/2, b.pos.Y+h, dxlib.GetColor(255, 255, 255), true)

		// 範囲
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		dxlib.DrawCircle(b.pos.X, b.pos.Y, bombBoulderRange, dxlib.GetColor(255, 255, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}
}

func (b *bombBoulder) Update() bool {
	b.count++
	return false
}
