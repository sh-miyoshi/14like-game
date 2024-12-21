package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Skill interface {
	Init(manager models.Manager)
	End()
	Draw()
	Update() bool
}
