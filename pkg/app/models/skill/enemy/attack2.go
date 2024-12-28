package skill

import (
	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

const (
	attack2CastTime = 2 * 60
)

type Attack2 struct {
	count   int
	ownerID string
	manager models.Manager

	rotateBase   point.Point
	viewStartPos point.Point
	width        int
	length       int
	angle        float64
}

func (a *Attack2) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *Attack2) End() {
}

func (a *Attack2) Draw() {
	// 詠唱バー
	castTime := attack2CastTime - a.count
	if a.count != 0 && castTime > 0 {
		size := 200
		objs := a.manager.GetObjectParams(&models.ObjectFilter{ID: a.ownerID})
		if len(objs) == 0 {
			return
		}
		px := objs[0].Pos.X - size/2
		py := objs[0].Pos.Y + 50
		dxlib.DrawBox(px, py, px+size, py+20, dxlib.GetColor(255, 255, 255), false)
		castSize := size * castTime / attack2CastTime
		dxlib.DrawBox(px, py, px+castSize, py+20, dxlib.GetColor(255, 255, 255), true)
		dxlib.DrawFormatString(px, py+25, 0xffffff, "範囲攻撃2")
	}

	// 範囲
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_ALPHA, 64)
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
	dxlib.SetDrawBlendMode(dxlib.DX_BLENDMODE_NOBLEND, 0)
}

func (a *Attack2) Update() bool {
	if a.count == 0 {
		objs := a.manager.GetObjectParams(&models.ObjectFilter{ID: a.ownerID})
		a.setParams(objs[0].Pos, objs[0].Pos, 100, config.ScreenSizeX, math.Pi/2)
	}

	// 詠唱
	if a.count >= attack2CastTime {
		a.manager.AddDamage(models.Damage{
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
		return true
	}

	a.count++
	return false
}

func (a *Attack2) setParams(rotBase, viewStart point.Point, width, length int, angle float64) {
	a.rotateBase = rotBase
	a.viewStartPos = viewStart
	a.width = width
	a.length = length
	a.angle = angle
}
