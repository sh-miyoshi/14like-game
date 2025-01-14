package skill

import (
	"fmt"
	"math/rand"

	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	WaveGunAttackLeft int = iota
	WaveGunAttackRight
)

const (
	waveGunCastTime = 180
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
	const interval = 60
	startTimes := [10]int{0, 0, 0, 2, 2, 1, 1, 3, 3, 3}
	isLeftFirst := rand.Intn(2) == 0
	if isLeftFirst {
		tmp := startTimes[0:5]
		math.Shuffle(tmp)
		for i := 0; i < 5; i++ {
			startTimes[i] = tmp[i]
		}
	} else {
		tmp := startTimes[5:10]
		math.Shuffle(tmp)
		for i := 0; i < 5; i++ {
			startTimes[i+5] = tmp[i]
		}
	}

	// 左半分
	for i := 0; i < 5; i++ {
		pos := point.Point{X: i * 65, Y: 0}
		pos = math.Rotate(point.Point{X: 65 * 5, Y: 0}, pos, math.Pi/4)

		id := a.manager.AddObject(models.ObjectInstWaveGunAttacker, &WaveGunAttackerParam{
			Pos: point.Point{
				X: pos.X + 60,
				Y: pos.Y + config.ScreenSizeY - 15,
			},
			Direct:    WaveGunAttackRight,
			StartTime: startTimes[i]*interval + waveGunCastTime,
		})
		a.attackIDs = fmt.Sprintf("%s,%s", a.attackIDs, id)
	}
	// 右半分
	for i := 0; i < 5; i++ {
		pos := point.Point{X: i * 65, Y: 0}
		pos = math.Rotate(point.Point{X: 0, Y: 0}, pos, -math.Pi/4)

		id := a.manager.AddObject(models.ObjectInstWaveGunAttacker, &WaveGunAttackerParam{
			Pos: point.Point{
				X: pos.X + 50 + config.ScreenSizeX/2,
				Y: pos.Y + config.ScreenSizeY - 50 - 15,
			},
			Direct:    WaveGunAttackLeft,
			StartTime: startTimes[i+5]*interval + waveGunCastTime,
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
	return a.count >= waveGunCastTime
}

func (a *WaveGun) GetCount() int {
	return a.count
}

func (a *WaveGun) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: waveGunCastTime,
		Name:     "斉射式波動砲",
	}
}
