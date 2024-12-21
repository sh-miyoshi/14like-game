package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Skill interface {
	Init()
	End()
	Draw()
	Update(manager models.Manager) bool
}
