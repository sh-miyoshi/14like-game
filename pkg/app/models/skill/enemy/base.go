package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Skill interface {
	Init(manager models.Manager, ownerID string)
	End()
	Draw()
	Update() bool
}
