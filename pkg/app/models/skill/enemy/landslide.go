package skill

import (
	orgmath "math"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
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
	// 詠唱バー
	castTime := landslideCastTime - a.count
	if a.count != 0 && castTime > 0 {
		size := 200
		objs := a.manager.GetObjectParams(&models.ObjectFilter{ID: a.ownerID})
		if len(objs) == 0 {
			return
		}
		px := objs[0].Pos.X - size/2
		py := objs[0].Pos.Y + 50
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / landslideCastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "範囲攻撃2")
	}

	// 範囲
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
	for _, a := range a.attack {
		a.Draw()
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
		a.attack[1].SetParams(myPos, view, w, config.ScreenSizeX, angle+math.Pi*4/6)
		a.attack[2].SetParams(myPos, view, w, config.ScreenSizeX, angle+math.Pi*2/6)
	}

	// 詠唱
	if a.count >= landslideCastTime {
		for _, atk := range a.attack {
			atk.AddDamage(a.manager)
		}
		return true
	}

	a.count++
	return false
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

func (a *landslideAttack) AddDamage(manager models.Manager) {
	manager.AddDamage(models.Damage{
		ID:         uuid.New().String(),
		Power:      100,
		DamageType: models.TypeAreaRect,
		RectPos: [2]point.Point{
			a.viewStartPos,
			{X: a.viewStartPos.X + a.width, Y: a.viewStartPos.Y + a.length},
		},
		RotateBase:  a.rotateBase,
		RotateAngle: a.angle,
	})
}
