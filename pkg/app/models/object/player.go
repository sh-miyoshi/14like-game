package object

import (
	"fmt"
	orgmath "math"

	"github.com/google/uuid"
	"github.com/sh-miyoshi/14like-game/pkg/app/config"
	"github.com/sh-miyoshi/14like-game/pkg/app/models"
	"github.com/sh-miyoshi/14like-game/pkg/dxlib"
	"github.com/sh-miyoshi/14like-game/pkg/inputs"
	"github.com/sh-miyoshi/14like-game/pkg/logger"
	"github.com/sh-miyoshi/14like-game/pkg/sound"
	"github.com/sh-miyoshi/14like-game/pkg/utils/math"
	"github.com/sh-miyoshi/14like-game/pkg/utils/point"
)

type Player struct {
	id      string
	pos     point.Point
	buffs   []models.Buff
	manager models.Manager
	hits    int
	direct  float64
}

func (p *Player) Init(manager models.Manager) {
	p.id = uuid.New().String()
	p.manager = manager
	p.pos.X = config.ScreenSizeX / 2
	p.pos.Y = config.ScreenSizeY * 3 / 4
	p.buffs = make([]models.Buff, 0)
	p.direct = math.Pi
}

func (p *Player) End() {
}

func (p *Player) Draw() {
	dxlib.DrawCircle(p.pos.X, p.pos.Y, config.PlayerHitRange, dxlib.GetColor(255, 255, 255), false)

	dxlib.DrawFormatString(10, 30, 0xFFFFFF, "ダメージを食らった回数: %d", p.hits)

	// バフ・デバフ
	for i, b := range p.buffs {
		icon := b.GetIcon()
		px := p.pos.X + config.PlayerHitRange/2 + 20
		py := p.pos.Y - config.PlayerHitRange/2 - 40
		dxlib.DrawGraph(px, py+i*32, icon, true)
		c := b.GetCount()/60 + 1
		dxlib.DrawStringToHandle(px+8, py+i*32+28, 0xffffff, config.SkillNumberFontHandle, fmt.Sprintf("%2d", c))

		// WIP: stack count
	}

	s := 10
	p1 := math.Rotate(p.pos, point.Point{X: p.pos.X - s/2, Y: p.pos.Y}, p.direct)
	p2 := math.Rotate(p.pos, point.Point{X: p.pos.X + s/2, Y: p.pos.Y}, p.direct)
	p3 := math.Rotate(p.pos, point.Point{X: p.pos.X, Y: p.pos.Y + s}, p.direct)
	dxlib.DrawTriangle(p1.X, p1.Y, p2.X, p2.Y, p3.X, p3.Y, dxlib.GetColor(255, 255, 255), true)
}

func (p *Player) Update() bool {
	for i := 0; i < len(p.buffs); i++ {
		if p.buffs[i].Update() {
			p.buffs[i].End()
			p.buffs = append(p.buffs[:i], p.buffs[i+1:]...)
			i--
		}
	}

	// Move
	spd := 3

	moveLR := 0
	moveUD := 0
	if inputs.CheckKey(inputs.KeyUp) > 0 {
		if p.pos.Y > config.PlayerHitRange {
			moveUD = -spd
		}
	} else if inputs.CheckKey(inputs.KeyDown) > 0 {
		if p.pos.Y < config.ScreenSizeY-config.PlayerHitRange {
			moveUD = spd
		}
	}

	if inputs.CheckKey(inputs.KeyRight) > 0 {
		if p.pos.X < config.ScreenSizeX-config.PlayerHitRange {
			moveLR = spd
		}
	} else if inputs.CheckKey(inputs.KeyLeft) > 0 {
		if p.pos.X > config.PlayerHitRange {
			moveLR = -spd
		}
	}
	if moveLR != 0 && moveUD != 0 {
		// NOTE: 本来は√2で割るべきだが、見栄え的な観点で1.2にしている
		moveLR = int(float64(moveLR) / 1.2)
		moveUD = int(float64(moveUD) / 1.2)
	}
	p.setDirect(moveLR, moveUD)
	p.pos.X += moveLR
	p.pos.Y += moveUD

	return false
}

func (p *Player) GetParam() models.ObjectParam {
	return models.ObjectParam{
		ID:       p.id,
		Pos:      p.pos,
		IsPlayer: true,
		Direct:   p.direct,
	}
}

func (p *Player) HandleDamage(dm models.Damage) {
	logger.Debug("NonAttackPlayer got damage %+v", dm)
	if dm.Power > 0 {
		sound.On(sound.SEFailed)
		p.hits++
		// WIP: リファクタリング
		p.manager.SetResult(models.ResultInfo{Hits: p.hits})
	}
	for _, b := range dm.Buffs {
		b.Init(p.manager, p.id)
		p.buffs = append(p.buffs, b)
	}
	if dm.Push != nil {
		if !dm.Push.IsBack && point.Distance2(p.pos, dm.Push.At) < int(dm.Push.Length*dm.Push.Length) {
			p.pos = dm.Push.At
		} else {
			rad := orgmath.Atan2(float64(p.pos.Y-dm.Push.At.Y), float64(p.pos.X-dm.Push.At.X))
			if dm.Push.IsBack {
				p.pos.X += int(dm.Push.Length * orgmath.Cos(rad))
				p.pos.Y += int(dm.Push.Length * orgmath.Sin(rad))
			} else {
				p.pos.X -= int(dm.Push.Length * orgmath.Cos(rad))
				p.pos.Y -= int(dm.Push.Length * orgmath.Sin(rad))
			}
		}
	}
}

func (p *Player) setDirect(moveLR, moveUD int) {
	if moveLR == 0 && moveUD == 0 {
		return
	}

	if moveLR == 0 {
		if moveUD > 0 {
			p.direct = 0
		} else {
			p.direct = math.Pi
		}
		return
	}

	if moveUD == 0 {
		if moveLR > 0 {
			p.direct = math.Pi * 3 / 2
		} else {
			p.direct = math.Pi / 2
		}
		return
	}

	if moveLR < 0 {
		if moveUD > 0 {
			p.direct = math.Pi / 4
		} else {
			p.direct = math.Pi * 3 / 4
		}
	} else {
		if moveUD > 0 {
			p.direct = math.Pi * 7 / 4
		} else {
			p.direct = math.Pi * 5 / 4
		}
	}
}
