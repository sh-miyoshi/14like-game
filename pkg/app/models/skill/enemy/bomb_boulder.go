package skill

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	bombBoulderInitDelay = 45
	bombBoulderCastTime  = 180
	bombBoulderRange     = 180
)

type bombBoulder struct {
	pos       point.Point
	count     int
	startTime int
	manager   models.Manager
	isEnd     bool
}

type BombBoulderMgr struct {
	ownerID string
	manager models.Manager
	attacks [9]bombBoulder
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
			a.attacks[y*3+x] = bombBoulder{
				pos: point.Point{
					X: (config.ScreenSizeX/4+40)*(x+1) - 40,
					Y: config.ScreenSizeY/4*(yindex[y]+1) - 20,
				},
				startTime: bombBoulderInitDelay * y,
				manager:   manager,
			}
		}
	}
	logger.Debug("bomb boulders: %+v", a.attacks)
}

func (a *BombBoulderMgr) End() {
}

func (a *BombBoulderMgr) Draw() {
	for _, b := range a.attacks {
		if !b.IsEnd() {
			b.Draw()
		}
	}
}

func (a *BombBoulderMgr) Update() bool {
	end := true
	for i := range a.attacks {
		if !a.attacks[i].IsEnd() {
			a.attacks[i].Update()
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
	}

	if b.count >= b.startTime+bombBoulderCastTime-20 {
		// 範囲
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
		dxlib.DrawCircle(b.pos.X, b.pos.Y, bombBoulderRange, dxlib.GetColor(255, 255, 0), true)
		dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
	}
}

func (b *bombBoulder) Update() {
	b.count++
	if b.count == b.startTime+bombBoulderCastTime {
		b.manager.AddDamage(models.Damage{
			ID:         uuid.New().String(),
			Power:      120,
			DamageType: models.DamageTypeAreaCircle,
			CenterPos:  b.pos,
			Range:      bombBoulderRange,
		})
		b.isEnd = true
	}
}

func (b *bombBoulder) IsEnd() bool {
	return b.isEnd
}
