package skill

type Param struct {
	CastTime   int
	RecastTime int
	Power      int
	Range      int
}

type Skill interface {
	Init()
	End()

	GetParam() Param
	GetIcon() int
}
