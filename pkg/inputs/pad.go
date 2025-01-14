package inputs

import "github.com/sh-miyoshi/14like-game/pkg/dxlib"

const padButtonNum = 28

type pad struct {
	keyBind  [keyMax]int
	padState [padButtonNum]int
}

func (p *pad) Init() error {
	p.keyBind[KeyEnter] = 6
	p.keyBind[KeyLeft] = 2
	p.keyBind[KeyRight] = 3
	p.keyBind[KeyUp] = 4
	p.keyBind[KeyDown] = 1

	return nil
}

func (p *pad) KeyStateUpdate() {
	state := dxlib.GetJoypadInputState(dxlib.DX_INPUT_PAD1 | dxlib.DX_INPUT_KEY)
	for i := 0; i < padButtonNum; i++ {
		if state&(1<<i) != 0 {
			p.padState[i]++
		} else {
			p.padState[i] = 0
		}
	}
}

func (p *pad) CheckKey(key KeyType) int {
	// 別ボタンはなし
	if key == KeyAnotherLeft || key == KeyAnotherRight || key == KeyAnotherUp || key == KeyAnotherDown {
		return 0
	}

	return p.padState[p.keyBind[key]-1]
}
