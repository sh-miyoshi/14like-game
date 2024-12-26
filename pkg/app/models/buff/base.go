package buff

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Buff interface {
	Init(manager models.Manager, ownerID string)
	End()
	Update() bool
	GetIcon() int
	GetCount() int
}
