package skill

import (
	"fmt"

	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	WaveGunAttackLeft int = iota
	WaveGunAttackRight
)

type WaveGunAttackerParam struct {
	StartTime int
	Pos       point.Point
	Direct    int
}

type WaveGun struct {
	count     int
	ownerID   string
	manager   models.Manager
	attackIDs string
}

func (a *WaveGun) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID

	// WIP: ランダム化
	startTimes := [10]int{0, 0, 0, 120, 120, 60, 60, 180, 180, 180}
	const initDelay = 60

	// 左半分
	for i := 0; i < 5; i++ {
		pos := point.Point{X: i * 65, Y: 0}
		pos = math.Rotate(point.Point{X: 65 * 5, Y: 0}, pos, math.Pi/4)

		id := a.manager.AddObject(models.ObjectTypeWaveGunAttacker, &WaveGunAttackerParam{
			Pos: point.Point{
				X: pos.X + 60,
				Y: pos.Y + config.ScreenSizeY - 15,
			},
			Direct:    WaveGunAttackRight,
			StartTime: startTimes[i] + initDelay,
		})
		a.attackIDs = fmt.Sprintf("%s,%s", a.attackIDs, id)
	}
	// 右半分
	for i := 0; i < 5; i++ {
		pos := point.Point{X: i * 65, Y: 0}
		pos = math.Rotate(point.Point{X: 0, Y: 0}, pos, -math.Pi/4)

		id := a.manager.AddObject(models.ObjectTypeWaveGunAttacker, &WaveGunAttackerParam{
			Pos: point.Point{
				X: pos.X + 50 + config.ScreenSizeX/2,
				Y: pos.Y + config.ScreenSizeY - 50 - 15,
			},
			Direct:    WaveGunAttackLeft,
			StartTime: startTimes[i+5] + initDelay,
		})
		a.attackIDs = fmt.Sprintf("%s,%s", a.attackIDs, id)
	}
}

func (a *WaveGun) End() {
}

func (a *WaveGun) Draw() {
}

func (a *WaveGun) Update() bool {
	a.count++
	return false
}

func (a *WaveGun) GetCount() int {
	return a.count
}

func (a *WaveGun) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: 10,
		Name:     "Attack",
	}
}
