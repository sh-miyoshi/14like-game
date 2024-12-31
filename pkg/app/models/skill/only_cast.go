package skill

import "github.com/sh-miyoshi/14like-game/pkg/app/models"

type OnlyCast struct {
	CastTime int
	Name     string

	count   int
	ownerID string
	manager models.Manager
}

func (a *OnlyCast) Init(manager models.Manager, ownerID string) {
	a.manager = manager
	a.ownerID = ownerID
}

func (a *OnlyCast) End() {
}

func (a *OnlyCast) Draw() {
}

func (a *OnlyCast) Update() bool {
	a.count++
	return a.count >= a.CastTime
}

func (a *OnlyCast) GetCount() int {
	return a.count
}

func (a *OnlyCast) GetParam() models.SkillParam {
	return models.SkillParam{
		CastTime: a.CastTime,
		Name:     a.Name,
	}
}
