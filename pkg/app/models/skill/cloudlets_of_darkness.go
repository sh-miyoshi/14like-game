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
	CloudletsOfDarknessAttackLeft int = iota
	CloudletsOfDarknessAttackRight
)

const (
	cloudletsOfDarknessCastTime = 180
)

type CloudletsOfDarknessAttackerParam struct {
	StartTime int
	Pos       point.Point
	Direct    int
}

type CloudletsOfDarkness struct {
	count     int
	ownerID   string
	manager   models.Manager
	attackIDs string
}

func (a *CloudletsOfDarkness) Init(manager models.Manager, ownerID string) {
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

		id := a.manager.AddObject(models.ObjectInstCloudletsOfDarknessAttacker, &CloudletsOfDarknessAttackerParam{
			Pos: point.Point{
				X: pos.X + 60,
				Y: pos.Y + config.ScreenSizeY - 15,
			},
			Direct:    CloudletsOfDarknessAttackRight,
			StartTime: startTimes[i]*interval + cloudletsOfDarknessCastTime,
		})
		a.attackIDs = fmt.Sprintf("%s,%s", a.attackIDs, id)
	}
	// 右半分
	for i := 0; i < 5; i++ {
		pos := point.Point{X: i * 65, Y: 0}
		pos = math.Rotate(point.Point{X: 0, Y: 0}, pos, -math.Pi/4)

		id := a.manager.AddObject(models.ObjectInstCloudletsOfDarknessAttacker, &CloudletsOfDarknessAttackerParam{
			Pos: point.Point{
				X: pos.X + 50 + config.ScreenSizeX/2,
				Y: pos.Y + config.ScreenSizeY - 50 - 15,
			},
			Direct:    CloudletsOfDarknessAttackLeft,
			StartTime: startTimes[i+5]*interval + cloudletsOfDarknessCastTime,
		})
		a.attackIDs = fmt.Sprintf("%s,%s", a.attackIDs, id)
	}
}

func (a *CloudletsOfDarkness) End() {
}

func (a *CloudletsOfDarkness) Draw() {
}

func (a *CloudletsOfDarkness) Update() bool {
	a.count++
	return a.count >= cloudletsOfDarknessCastTime
}

func (a *CloudletsOfDarkness) GetCount() int {
	return a.count
}

func (a *CloudletsOfDarkness) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: cloudletsOfDarknessCastTime,
		Name:     "斉射式波動砲",
	}
}
