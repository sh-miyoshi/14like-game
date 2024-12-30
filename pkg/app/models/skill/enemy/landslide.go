package skill

import (
	orgmath "math"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/app/models/buff"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	landslideCastTime = 2 * 60
)

type landslideAttack struct {
	rotateBase   point.Point
	viewStartPos point.Point
	width        int
	length       int
	angle        float64
}

type LandSlide struct {
	AttackNum int

	count   int
	ownerID string
	manager models.Manager
	attack  [3]landslideAttack
}

func (a *LandSlide) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *LandSlide) End() {
}

func (a *LandSlide) Draw() {
	// 範囲
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	for i := 0; i < a.AttackNum; i++ {
		a.attack[i].Draw()
	}
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *LandSlide) Update() bool {
	if a.count == 0 {
		objs := a.manager.GetObjectParams(nil)
		var myPos, targetPos point.Point
		for _, obj := range objs {
			if obj.ID == a.ownerID {
				myPos = obj.Pos
			} else {
				targetPos = obj.Pos
			}
		}
		angle := orgmath.Atan2(float64(myPos.Y-targetPos.Y), float64(myPos.X-targetPos.X))
		w := 150
		view := point.Point{X: myPos.X - w/2, Y: myPos.Y}
		a.attack[0].SetParams(myPos, view, w, config.ScreenSizeX, angle+math.Pi/2)
		if a.AttackNum >= 3 {
			a.attack[1].SetParams(myPos, view, w, config.ScreenSizeX, angle+math.Pi*4/6)
			a.attack[2].SetParams(myPos, view, w, config.ScreenSizeX, angle+math.Pi*2/6)
		}
	}

	// 詠唱
	if a.count >= landslideCastTime {
		for i := 0; i < a.AttackNum; i++ {
			a.attack[i].AddDamage(a.manager, a.ownerID)
		}
		return true
	}

	a.count++
	return false
}

func (a *LandSlide) GetCount() int {
	return a.count
}

func (a *LandSlide) GetParam() models.EnemySkillParam {
	return models.EnemySkillParam{
		CastTime: landslideCastTime,
		Name:     "ランドスライド",
	}
}

func (a *landslideAttack) SetParams(rotBase, viewStart point.Point, width, length int, angle float64) {
	a.rotateBase = rotBase
	a.viewStartPos = viewStart
	a.width = width
	a.length = length
	a.angle = angle
}

func (a *landslideAttack) Draw() {
	b := a.rotateBase
	s := a.viewStartPos
	p1 := math.Rotate(b, s, a.angle)
	p2 := math.Rotate(b, point.Point{X: s.X + a.width, Y: s.Y}, a.angle)
	p3 := math.Rotate(b, point.Point{X: s.X + a.width, Y: s.Y + a.length}, a.angle)
	p4 := math.Rotate(b, point.Point{X: s.X, Y: s.Y + a.length}, a.angle)
	dxlib.DrawQuadrangle(
		p1.X, p1.Y,
		p2.X, p2.Y,
		p3.X, p3.Y,
		p4.X, p4.Y,
		dxlib.GetColor(255, 255, 0),
		true,
	)
}

func (a *landslideAttack) AddDamage(manager models.Manager, ownerID string) {
	manager.AddDamage(models.Damage{
		ID:         uuid.New().String(),
		Power:      100,
		DamageType: models.DamageTypeAreaRect,
		RectPos: [2]point.Point{
			a.viewStartPos,
			{X: a.viewStartPos.X + a.width, Y: a.viewStartPos.Y + a.length},
		},
		RotateBase:  a.rotateBase,
		RotateAngle: a.angle,
		Buffs:       []models.Buff{&buff.UpDamage{}},
	})
}
