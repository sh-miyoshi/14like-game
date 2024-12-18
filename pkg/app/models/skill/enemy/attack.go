package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type Attack struct {
}

func (a *Attack) Init() {
}

func (a *Attack) End() {
}

func (a *Attack) Exec(AddDamage func(models.Damage)) {
}

func (a *Attack) GetParam() Param {
	return Param{}
}
